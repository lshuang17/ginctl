// @Name: dao.go
// @Date: 2023-06-19
// @Author: ls

package demo

type IDemoDao interface {
}

type demoDao struct {
     
}

func NewDemoDao( ) IDemoDao {
    
    return &demoDao{
    }
}
