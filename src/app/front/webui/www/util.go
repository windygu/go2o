/**
 * Copyright 2013 @ ops Inc.
 * name :
 * author : jarryliu
 * date : 2014-02-03 23:18
 * description :
 * history :
 */
package www

import (
	"github.com/atnet/gof"
	"go2o/src/core/domain/interface/enum"
	"go2o/src/core/domain/interface/partner"
	"go2o/src/core/service/goclient"
	"html/template"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
)

// 跳转到登录页面
func RedirectLoginPage(w http.ResponseWriter, returnUrl string) {
	var header http.Header = w.Header()
	header.Add("Location", "/login?return_url="+url.QueryEscape(returnUrl))
	w.WriteHeader(302)
}

func GetSiteConf(w http.ResponseWriter, p *partner.ValuePartner) (bool, *partner.SiteConf) {
	siteConf, _ := goclient.Partner.GetSiteConf(p.Id, p.Secret)

	if siteConf == nil {
		w.Write([]byte("网站访问过程中出现了异常，请重试!"))
		return false, nil
	}

	if siteConf.State == enum.PARTNER_SITE_CLOSED {
		if strings.TrimSpace(siteConf.StateHtml) == "" {
			siteConf.StateHtml = "网站暂停访问，请联系商家：" + p.Tel
		}
		w.Write([]byte(siteConf.StateHtml))
		return false, siteConf
	}
	return true, siteConf
}

// 处理自定义错误
func handleCustomError(w http.ResponseWriter, ctx gof.App, err error) {
	if err != nil {
		ctx.Template().Execute(w, func(m *map[string]interface{}) {
			(*m)["error"] = err.Error()
			(*m)["statck"] = template.HTML(strings.Replace(string(debug.Stack()), "\n", "<br />", -1))
		},
			"views/web/www/error.html")
	}
}
