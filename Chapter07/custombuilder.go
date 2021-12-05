package main

import (
  "context"
  "fmt"
  "github.com/containers/buildah"
  "github.com/containers/buildah/define"
  "github.com/containers/storage/pkg/unshare"
  is "github.com/containers/image/v5/storage"
  "github.com/containers/image/v5/types"
  "github.com/containers/storage"
)

func main() {
  if buildah.InitReexec() {
    return
  }
  unshare.MaybeReexecUsingUserNamespace(false)

  buildStoreOptions, err := storage.DefaultStoreOptions(unshare.IsRootless(), unshare.GetRootlessUID())

  if err != nil {
    panic(err)
  }

  buildStore, err := storage.GetStore(buildStoreOptions)

  if err != nil {
    panic(err)
  }

  opts := buildah.BuilderOptions{
    FromImage:        "node:12-alpine",
    Isolation:        define.IsolationChroot,
    CommonBuildOpts:  &define.CommonBuildOptions{},
    ConfigureNetwork: define.NetworkDefault,
    SystemContext:    &types.SystemContext {},
  }

  builder, err := buildah.NewBuilder(context.TODO(), buildStore, opts)

  if err != nil {
    panic(err)
  }

  err = builder.Add("/home/node/", false, buildah.AddAndCopyOptions{}, "script.js")

  if err != nil {
    panic(err)
  }

  builder.SetCmd([]string{"node", "/home/node/script.js"})

  imageRef, err := is.Transport.ParseStoreReference(buildStore, "docker.io/myusername/my-image")

  if err != nil {
    panic(err)
  }

  imageId, _, _, err := builder.Commit(context.TODO(), imageRef, buildah.CommitOptions{})

  fmt.Printf("Image built! %s\n", imageId)
}
