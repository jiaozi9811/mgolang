```
func TestIndexRouter(t *testing.T){
      router:=gin.default()
      w:=httptest.NewRecorder()
      req,_:=http.NewRequest(http.MethodPost,"/",nil)
      router.ServeHTTP(w,req)
      assert.Equal(t,http.StatusOK,w.Code)
      assert.Equal(t,"hell gin post method",w.Body.String())
}
```
