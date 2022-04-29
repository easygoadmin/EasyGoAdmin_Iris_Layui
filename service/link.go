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

var Link = new(linkService)

type linkService struct{}

func (s *linkService) GetList(req dto.LinkPageReq) ([]vo.LinkInfoVo, int64, error) {
	// 实例化对象
	query := utils.XormDb.Where("mark=1")
	// 友链名称
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	// 友链类型
	if req.Type > 0 {
		query = query.Where("type=?", req.Type)
	}
	// 投放平台
	if req.Platform > 0 {
		query = query.Where("platform=?", req.Platform)
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 对象转换
	var list []model.Link
	count, err := query.FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}

	var result []vo.LinkInfoVo
	for _, v := range list {
		item := vo.LinkInfoVo{}
		item.Link = v
		// 友链图片
		if v.Image != "" {
			item.Image = utils.GetImageUrl(v.Image)
		}
		// 友链类型
		typeName, ok := constant.LINK_TYPE_LIST[v.Type]
		if ok {
			item.TypeName = typeName
		}
		// 友链形式
		formName, ok := constant.LINK_FORM_LIST[v.Form]
		if ok {
			item.FormName = formName
		}
		// 投放平台
		platformName, ok := constant.LINK_PLATFORM_LIST[v.Platform]
		if ok {
			item.PlatformName = platformName
		}
		result = append(result, item)
	}
	// 返回结果
	return result, count, nil
}

func (s *linkService) Add(req dto.LinkAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.Link
	entity.Name = req.Name
	entity.Type = gconv.Int(req.Type)
	entity.Url = req.Url
	entity.ItemId = gconv.Int(req.ItemId)
	entity.CateId = gconv.Int(req.CateId)
	entity.Platform = gconv.Int(req.Platform)
	entity.Form = gconv.Int(req.Form)
	entity.Status = gconv.Int(req.Status)
	entity.Sort = gconv.Int(req.Sort)
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	entity.Mark = 1

	// 图片处理
	if req.Image != "" {
		image, err := utils.SaveImage(req.Image, "link")
		if err != nil {
			return 0, err
		}
		entity.Image = image
	}

	// 插入数据
	return entity.Insert()
}

func (s *linkService) Update(req dto.LinkUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.Link{Id: gconv.Int(req.Id)}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置对象
	entity.Name = req.Name
	entity.Type = gconv.Int(req.Type)
	entity.Url = req.Url
	entity.ItemId = gconv.Int(req.ItemId)
	entity.CateId = gconv.Int(req.CateId)
	entity.Platform = gconv.Int(req.Platform)
	entity.Form = gconv.Int(req.Form)
	entity.Status = gconv.Int(req.Status)
	entity.Sort = gconv.Int(req.Sort)
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()

	// 图片处理
	if req.Image != "" {
		image, err := utils.SaveImage(req.Image, "link")
		if err != nil {
			return 0, err
		}
		entity.Image = image
	}

	// 更新记录
	return entity.Update()
}

func (s *linkService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.Link{Id: gconv.Int(ids)}
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
			entity := &model.Link{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *linkService) Status(req dto.LinkStatusReq, userId int) (int64, error) {
	// 查询记录是否存在
	info := &model.Link{Id: gconv.Int(req.Id)}
	has, err := info.Get()
	if err != nil || !has {
		return 0, errors.New("记录不存在")
	}

	// 设置状态
	entity := &model.Link{}
	entity.Id = info.Id
	entity.Status = gconv.Int(req.Status)
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	return entity.Update()
}
