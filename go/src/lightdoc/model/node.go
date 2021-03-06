package model

import (
	"io/ioutil"
	"lightdoc/config"
	"path"
	"strings"
)

type Node struct {
	//Node 为文件夹
	Nodes []Node
	//Path为相对路径
	Path, Name string
	//Name
	IsDir bool
}

func InitTree(path string) (tree *Node) {
	tree = new(Node)
	tree.Filelist(path)
	return
}
func (tree *Node) Filelist(p string) {
	if tree.Nodes == nil {
		tree.Nodes = make([]Node, 0)
	}
	files, _ := ioutil.ReadDir(p)
	for _, f := range files {
		if !f.IsDir() {
			sname := path.Ext(f.Name())
			switch sname {
			case ".md":
			case ".pdf":
				{
				}
			default:
				continue
			}
		} else {
			sname := path.Ext(f.Name())
			switch sname {
			case ".git":
			case ".vscode":
			case ".idea":
				{
					continue
				}
			default:

			}
		}
		childdir := p + "/" + f.Name()
		childnode := new(Node)
		childnode.IsDir = f.IsDir()
		childnode.Path = childdir[len(config.Root):]
		fname := path.Base(f.Name())
		sname := path.Ext(f.Name())
		filenameOnly := strings.TrimSuffix(fname, sname)
		childnode.Name = filenameOnly
		childnode.Filelist(childdir)

		tree.Nodes = append(tree.Nodes, *childnode)

	}

}
