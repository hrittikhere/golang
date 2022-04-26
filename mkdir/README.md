Works like `mkdir -n foldername`

```bash
@hrittikhere ➜ /workspaces/golang/mkdir (main ✗) $ ls
go.mod  main.go  README.md

@hrittikhere ➜ /workspaces/golang/mkdir (main ✗) $ go run main.go -n New
New Created 

@hrittikhere ➜ /workspaces/golang/mkdir (main ✗) $ ls
go.mod  main.go  New  README.md
```