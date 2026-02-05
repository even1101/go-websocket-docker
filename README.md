docker run --rm -it --name go-demo -v PWD:/go/src/example.com/demo -p 8000:8080 golang

> in docker
1.go mod download 

2.go run cmd/web/*.go




docker doc ref: https://segmentfault.com/a/1190000021660020

Local StringVar text := {DetailTable.L_SubjectName};
Local NumberVar indentCount := {DetailTable.L_Level} * 2;
Local StringVar indent := ReplicateString("　", indentCount);  //全形空格
Local NumberVar maxLen := 25 - indentCount;  //每行最多字數
Local StringVar result := "";
Local NumberVar pos := 1;

While pos <= Len(text) Do (
    result := result + indent + Mid(text, pos, maxLen) + Chr(13) + Chr(10);
    pos := pos + maxLen
);

Left(result, Len(result) - 2)  //移除最後的換行
