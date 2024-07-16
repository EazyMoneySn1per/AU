export function getState() {
    return wx.getStorageSync('state')
}

export function setState(State) {
    return wx.setStorageSync('state', State)
}

export function removeState() {
    return wx.removeStorageSync('state')
}