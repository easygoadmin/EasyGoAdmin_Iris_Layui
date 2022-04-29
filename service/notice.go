package service

import (
	"easygoadmin/constant"
	"easygoadmin/dto"
	"easygoadmin/model"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"easygoadmin/vo"
	"errors"
	"strconv"
	"strings"
	"time"
)

var Notice = new(noticeService)

type noticeService struct{}

func (s *noticeService) GetList(req dto.NoticePageReq) ([]vo.NoticeInfoVo, int64, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 查询条件
	// 通知标题
	if req.Title != "" {
		query = query.Where("title like ?", "%"+req.Title+"%")
	}
	// 通知来源
	if req.Source > 0 {
		query = query.Where("source=?", req.Source)
	}
	// 排序
	query = query.OrderBy("id desc")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 对象转换
	var list []model.Notice
	count, err := query.FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}

	// 数据处理
	var result []vo.NoticeInfoVo
	for _, v := range list {
		item := vo.NoticeInfoVo{}
		item.Notice = v
		item.SourceName = constant.NOTICE_SOURCE_LIST[v.Source]
		result = append(result, item)
	}
	return result, count, nil
}

func (s *noticeService) Add(req dto.NoticeAddReq, userId int) (int64, error) {
	// 富文本处理
	content := utils.SaveImageContent(req.Content, req.Title, "notice")
	// 实例化对象
	var entity model.Notice
	entity.Title = req.Title
	entity.Content = content
	entity.IsTop = req.IsTop
	entity.Source = req.Source
	entity.Status = req.Status
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *noticeService) Update(req dto.NoticeUpdateReq, userId int) (int64, error) {
	// 查询记录
	// 查询记录
	entity := &model.Notice{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 富文本处理
	content := utils.SaveImageContent(req.Content, req.Title, "notice")

	// 设置参数
	entity.Title = req.Title
	entity.Content = content
	entity.IsTop = req.IsTop
	entity.Source = req.Source
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	// 更新记录
	return entity.Update()
}

func (s *noticeService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Notice{Id: gconv.Int(ids)}
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
			entity := &model.Notice{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *noticeService) Status(req dto.NoticeStatusReq, userId int) (int64, error) {
	// 查询记录是否存在
	info := &model.Notice{Id: req.Id}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity := &model.Notice{}
	entity.Id = info.Id
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	return entity.Update()
}
