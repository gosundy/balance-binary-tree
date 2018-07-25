package main

import "fmt"

type TreeNode struct {
	data int
	leftChild *TreeNode
	rightChild *TreeNode
	leftLevel int
	rightLevel int
}
func main(){
	//data:=[]int{4,5,3,2,1}
	data:=[]int{16,3,7,11,9,26,18,14,15}
	root:=createBalancedTree(data)
	zhongXuBianLiTree(root)
}
func createBalancedTree(data[]int)*TreeNode{
	if len(data)==0{
		return nil
	}
	root:=&TreeNode{data:data[0],leftLevel:0,rightLevel:0}
	for _,a:=range data[1:]{
		tmpRoot:=insertTree(root,a,nil,false)
		if tmpRoot!=nil{
			root=tmpRoot
		}
	}
	return root
}
//插入一个节点后，检查平衡的节点，然后调整该不平衡节点到平衡状态，通过递归回溯更新level
func insertTree(node *TreeNode,data int,parent *TreeNode,nodeIsLeft bool)(root *TreeNode){
	if node==nil{
		return
	}
	root=nil
	if data<node.data{
		insertTree(node.leftChild,data,node,true)
		if node.leftChild==nil{
			node.leftChild=&TreeNode{data:data,leftLevel:0,rightLevel:0}
		}
		//如果不平衡则调整
		node.leftLevel=max(node.leftChild.leftLevel,node.leftChild.rightLevel)+1
		if abs(node.leftLevel,node.rightLevel)>1{
			root=tiaoZhengShu(node,parent,nodeIsLeft)
		}

	}else{
		insertTree(node.rightChild,data,node,false)
		if node.rightChild==nil{
			node.rightChild=&TreeNode{data:data,leftLevel:0,rightLevel:0}
		}
		node.rightLevel=max(node.rightChild.leftLevel,node.rightChild.rightLevel)+1
		//如果不平衡则调整
		if abs(node.leftLevel,node.rightLevel)>1{
			root=tiaoZhengShu(node,parent,nodeIsLeft)
		}


	}
	return root
}
func tiaoZhengShu(node,nodeParent *TreeNode,nodeIsLeft bool)(root *TreeNode){
	root=nil
	if (node.leftLevel-node.rightLevel)>0{
		//LL调整
		if(node.leftChild.leftLevel-node.leftChild.rightLevel)>0{
			//说明node节点是root节点
			kNode:=node.leftChild
			if nodeParent==nil{
				//data:=[]int{5,3,6,2,4,1}
				root=kNode
			}else{
				if nodeIsLeft{
					nodeParent.leftChild=kNode
				}else{
					nodeParent.rightChild=kNode
				}
			}
			tmpRight:=kNode.rightChild
			kNode.rightLevel=0
			node.leftLevel=0
			node.leftChild=nil
			kNode.rightChild=node

			if tmpRight!=nil{
				//添加新节点更新level
				updateTreeLevel(node,tmpRight)
			}
			kNode.rightLevel=max(node.leftLevel,node.rightLevel)+1
			return root

		}else{
			//LR调整
			kNode:=node.leftChild.rightChild
			if nodeParent==nil{
				//data:=[]int{5,3,6,2,4,1}
				root=kNode
			}else{
				if nodeIsLeft{
					nodeParent.leftChild=kNode
				}else{
					nodeParent.rightChild=kNode
				}
			}
			tmpLeft:=kNode.leftChild
			tmpRight:=kNode.rightChild
			kNode.rightLevel=0
			kNode.leftLevel=0
			kNode.leftChild=node.leftChild
			kNode.rightChild=node
			kNode.leftChild.rightChild=nil
			kNode.leftChild.rightLevel=0
			kNode.rightChild.leftChild=nil
			kNode.rightChild.leftLevel=0


			if tmpLeft!=nil{
				//添加新节点更新level
				updateTreeLevel(kNode.leftChild,tmpLeft)
			}
			kNode.leftLevel=max(kNode.leftChild.leftLevel,kNode.leftChild.rightLevel)+1
			if tmpRight!=nil{
				//添加新节点更新level
				updateTreeLevel(kNode.rightChild,tmpRight)
			}
			kNode.rightLevel=max(kNode.rightChild.leftLevel,kNode.rightChild.rightLevel)+1

			return root
		}
	}else{
		//RR调整
		if(node.rightChild.leftLevel-node.rightChild.rightLevel)<0{
			//说明node节点是root节点
			kNode:=node.rightChild
			if nodeParent==nil{
				//data:=[]int{5,3,6,2,4,1}
				root=kNode
			}else{
				if nodeIsLeft{
					nodeParent.leftChild=kNode
				}else{
					nodeParent.rightChild=kNode
				}
			}
			tmpLeft:=kNode.leftChild
			kNode.leftLevel=0
			node.rightLevel=0
			node.rightChild=nil
			kNode.leftChild=node

			if tmpLeft!=nil{
				//添加新节点更新level
				updateTreeLevel(node,tmpLeft)
			}
			kNode.leftLevel=max(node.leftLevel,node.rightLevel)+1
			return root

		}else{
			//RL调整
			kNode:=node.rightChild.leftChild
			if nodeParent==nil{
				//data:=[]int{5,3,6,2,4,1}
				root=kNode
			}else{
				if nodeIsLeft{
					nodeParent.leftChild=kNode
				}else{
					nodeParent.rightChild=kNode
				}
			}
			tmpLeft:=kNode.leftChild
			tmpRight:=kNode.rightChild
			kNode.rightLevel=0
			kNode.leftLevel=0
			kNode.leftChild=node
			kNode.rightChild=node.rightChild
			kNode.leftChild.rightChild=nil
			kNode.leftChild.rightLevel=0
			kNode.rightChild.leftChild=nil
			kNode.rightChild.leftLevel=0


			if tmpLeft!=nil{
				//添加新节点更新level
				updateTreeLevel(kNode.leftChild,tmpLeft)
			}
			kNode.leftLevel=max(kNode.leftChild.leftLevel,kNode.leftChild.rightLevel)+1
			if tmpRight!=nil{
				//添加新节点更新level
				updateTreeLevel(kNode.rightChild,tmpRight)
			}
			kNode.rightLevel=max(kNode.rightChild.leftLevel,kNode.rightChild.rightLevel)+1

			return root
		}
	}
	return root
}
//插入更新，根据走的路径，逐级更新level
func updateTreeLevel(node *TreeNode,insertNode *TreeNode){
	if node==nil{
		return
	}
	if insertNode.data>=node.data{
		updateTreeLevel(node.rightChild,insertNode)
		if node.rightChild==nil{
			node.rightChild=insertNode
		}
		node.rightLevel=max(node.rightChild.leftLevel,node.rightChild.rightLevel)+1
	}else{
		updateTreeLevel(node.leftChild,insertNode)
		if node.leftChild==nil{
			node.leftChild=insertNode
		}
		node.leftLevel=max(node.leftChild.leftLevel,node.leftChild.rightLevel)+1
	}
}
func qianXuBianLiTree(node *TreeNode){
	if node==nil{
		return
	}
	fmt.Println(node.data," ",node.leftLevel," ",node.rightLevel)
	qianXuBianLiTree(node.leftChild)
	qianXuBianLiTree(node.rightChild)
}
func zhongXuBianLiTree(node *TreeNode){
	if node==nil{
		return
	}
	zhongXuBianLiTree(node.leftChild)
	fmt.Println(node.data," ",node.leftLevel," ",node.rightLevel)
	zhongXuBianLiTree(node.rightChild)
}
func houXuBianLiTree(node *TreeNode){
	if node==nil{
		return
	}
	houXuBianLiTree(node.leftChild)
	houXuBianLiTree(node.rightChild)
	fmt.Println(node.data," ",node.leftLevel," ",node.rightLevel)
}
func max(a,b int)int{

	if a>=b{

		return a
	}else{
		return b
	}
}
func abs(a,b int)int{
	if a>=b{
		return a-b
	}else{
		return b-a
	}
}


