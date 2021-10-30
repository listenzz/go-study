我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 Wrap 这个 error，抛给上层?

先说结论，要的。不仅要 wrap，而且要用 github.com/pkg/errors 进行 wrap。

堆栈信息对于调试问题至关重要，我们需要从根因开始记录堆栈信息，因此调用其他库（标准库、企业公共库、开源第三方库等）获取到错误时，要使用 `errors.Wrap` 添加堆栈信息。而 `sql.ErrNoRows` 这样的错误，是应用代码能用堆栈追溯到的源错误，即根因。

类似于 `sql.ErrNoRows` 这样的错误，我们应像如下处理：

```go
package dao

import (
	"github.com/pkg/errors"
)

func AddData() error {
    sql := "SELECT name FROM user"
    // 调用标准库，三方库，或私有库代码
	err := executeSQL(sql)
	if err != nil {
        // 有可能发生 sql.ErrNoRows，需要 wrap 根因，附加堆栈信息，同时记录是在执行什么 sql 语句发生的错误
		return errors.Wrap(err, sql)
	}
	return nil
}

```

如果已经对根因 wrap 过了，就不要再 wrap 了，可以使用 `errors.WithMessage` 提供额外的信息。

标准库的 wrap，即 `fmt.Errorf` 无法记录堆栈信息，故使用 github.com/pkg/errors。

```go
func main() {
	err := stdWrap()
	if err != nil {
        // 不会输出堆栈信息，无法有效调试
		fmt.Printf("%+v\n", err)
	}
}

func stdWrap() error {
	err := doStuff()
	if err != nil {
		return fmt.Errorf("std 干活的时候, %w", err)
	}
	return nil
}
```
