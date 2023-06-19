// @Name: service.go
// @Date: 2023-06-19
// @Author: ls

package demo

type IDemoService interface {
}

type demoService struct {
    dao IDemoDao
}

func NewDemoService(dao IDemoDao) IDemoService {
    return &demoService{
		dao: dao,
	}
}
