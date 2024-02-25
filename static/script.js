const dataObj={}

const subBtn=document.querySelector(".sub-btn")
subBtn.addEventListener("click",function(event){
    event.preventDefault()
    dataObj["title"]=document.querySelector(".input-title").value
    dataObj["body"]=document.querySelector(".input-body").value
    sendToServer(dataObj)
})
async function sendToServer(data){
    const options={
        method:"post",
        body:JSON.stringify(data)
    }
await fetch("http://localhost:5000/addnote",options)
}