4条注释：版权说明注释"//"、包说明注释"/*"(之后就是package)、函数说明注释"//"和最后添加的遗留问题说明"//"。

cd doctest
go doc doctest
godoc -http=:76
http://localhost:76/

若要将注释提取为文档，要遵守如下的基本规则。
 注释需要紧贴在对应的包声明和函数之前，不能有空行。
 注释如果要新起一个段落，应该用一个空白注释行隔开，因为直接换行书写会被认为是
正常的段内折行。
 开发者可以直接在代码内用// BUG(author): 的方式记录该代码片段中的遗留问题，这
些遗留问题也会被抽取到文档中。