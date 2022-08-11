/*
 * @Date: 2022-03-02 14:37:58
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:适合单服务器的缓存
 * @LastEditTime: 2022-03-28 09:39:37
 * @FilePath: \server\system\core\cache\cache.go
 */
package cache

import (
	"time"

	"chunDoc/system/util/convert"
	"github.com/patrickmn/go-cache"
)

var _cache *cache.Cache

func Instance() *cache.Cache {
	if _cache == nil {
		cacheInit()
	}
	return _cache
}

//默认缓存时间
func cacheInit() {
	_cache = cache.New(10*time.Minute, 20*time.Minute)
}

//设置缓存 -1 永久缓存   0 默默人
func Set(key interface{}, value interface{}, d time.Duration) {
	Instance().Set(convert.String(key), value, d)
}

//读取缓存
func Get(key interface{}) (interface{}, bool) {
	return Instance().Get(convert.String(key))
}

//删除缓存
func Delete(key interface{}) {
	Instance().Delete(convert.String(key))
}

//清空缓存
func Flush(key interface{}) {
	Instance().Flush()
}
