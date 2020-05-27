package main

import "fmt"

func main() {
    var (
        name string
        age  int
    )
    //n, _ := fmt.Sscanf("polaris 28", "%s%d", &name, &age)
    // 可以将"polaris 28"中的空格换成"\n"试试
    n, _ := fmt.Sscanf("polaris\n28", "%s%d", &name, &age)
    fmt.Println(n, name, age)
}
