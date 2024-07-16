const formatTime = date => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return `${[year, month, day].map(formatNumber).join('/')} ${[
    hour,
    minute,
    second,
  ]
    .map(formatNumber)
    .join(':')}`
}

const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : `0${n}`
}

// 第一个参数是需要进行防抖处理的函数，第二个参数是延迟时间，默认为1秒钟
function debounce(fn, delay = 1000) {
  let time = null
  return () => {
    time !== null && clearTimeout(time)
    time = setTimeout(fn, delay)
  }
}

module.exports = {
  formatTime,
  debounce,
}
