package service

import "gBlog/service/image_svc"

type ServiceGroup struct {
	ImageService image_svc.ImageService
}

var ServiceApp = new(ServiceGroup)
