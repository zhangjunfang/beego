package test

import (
	_ "api/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	//"fmt"
	"github.com/astaxie/beego"
	"github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	//fmt.Println(filepath.Join(file, "../"))
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../")))
	beego.TestBeegoInit(apppath)
}

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())

	convey.Convey("Subject: Test Station Endpoint\n", t, func() {
		convey.Convey("Status Code Should Be 200", func() {
			convey.So(w.Code, convey.ShouldEqual, 200)
		})
		convey.Convey("The Result Should Not Be Empty", func() {
			convey.So(w.Body.Len(), convey.ShouldBeGreaterThan, 0)
		})
	})
}
