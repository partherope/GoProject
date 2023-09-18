Email Address:hemoik.willot@hds.utc.fr

Les fonctions de chaque commande:
go version  - 显示Go版本

go build compilé et écris éxecutable dans ① là ou on exeécute la cmd ②là ou on a spécifié 一般是当前目录
go run  
go install compilé et écris exécutable dans .../go/bin  安装我们的程序到环境目录的bin下
gofmt / go fmt

go doc

Go的对象使用首字母大小写区分private和public参数和方法

全局constante变量必须是简单变量，不能是组合变量

### 切片是对原来数组的引用，修改它会修改原数组的值
func Fill(sl []int) {
	for i := 0; i < len(sl); i++ {
		sl[i] = rand.Int()
	}
}

