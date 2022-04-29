package service

import (
	"easygoadmin/dto"
	"easygoadmin/model"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"strconv"
	"strings"
	"time"
)

var Position = new(positionService)

type positionService struct{}

func (s *positionService) GetList(req dto.PositionPageReq) ([]model.Position, int64, error) {
	// 初始化查询实例
	query := utils.XormDb.Where("mark=1")
	// 岗位名称
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	// 排序
	query = query.OrderBy("sort")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询数据
	var list []model.Position
	count, err := query.FindAndCount(&list)
	return list, count, err
}

func (s *positionService) Add(req dto.PositionAddReq, userId int) (int64, error) {
	// 实例化模型
	var entity model.Position
	entity.Name = req.Name
	entity.Status = gconv.Int(req.Status)
	entity.Sort = gconv.Int(req.Sort)
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *positionService) Update(req dto.PositionUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Position{Id: gconv.Int(req.Id)}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}
	entity.Name = req.Name
	entity.Status = gconv.Int(req.Status)
	entity.Sort = gconv.Int(req.Sort)
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	// 更新数据
	return entity.Update()
}

func (s *positionService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := model.Position{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			id, _ := strconv.Atoi(v)
			entity := &model.Position{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *positionService) Status(req dto.PositionStatusReq, userId int) (int64, error) {
	// 查询记录
	info := &model.Position{Id: gconv.Int(req.Id)}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 更新状态
	entity := &model.Position{}
	entity.Id = gconv.Int(req.Id)
	entity.Status = gconv.Int(req.Status)
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	return entity.Update()
}
