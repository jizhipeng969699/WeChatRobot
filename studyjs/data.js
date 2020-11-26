"use strict"

let now = new Date()

console.log(now.toLocaleString())


let gotimestamp =new Date(2006,1,2,15,4,5)

console.log(gotimestamp.toLocaleString())



let unix = new Date(0)

console.log(unix)
console.log(parseInt(unix.getSeconds()))



let utc = new Date().toDateString()

console.log(utc)
console.log(typeof utc)


let mytime = new Date("2006-01-02 15:03:04")

console.log(mytime)



let gettime = new Date()
console.log(gettime.getTime())

console.log(Date(1606231370048).toLocaleString())



let setyear = new Date().setFullYear(1997).valueOf()
console.log(setyear)
    
    
    //get unix second

let unixsecond = new Date().getTime()
console.log(unixsecond)

let timefomart = new Date()
console.log(timefomart.toLocaleString())