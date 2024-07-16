const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  nickName: state => state.user.nickName,
  realName: state => state.user.realName,
  studentId: state => state.user.studentId,
  assName: state => state.user.assName,
  code: state => state.user.code
  // permission_routers: state => state.permission.routes// 路由列表
}
export default getters
