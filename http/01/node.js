const express = require('express'),
    app = express();


app.listen(4000, () => console.log("http://localhost:4000"))
app.get("/", (req, res) => {
    res.send("hello")
})
app.use("/post",(req,res)=>{
    res.send("post请求成功")
})