import { api } from './axios'
window.api = api

export const liveApi = {
  login({
    accountId,
    password
  } = {}) {
    return api.post('/console/v1/userLogin', {
      accountId,
      password
    })
  },
  getUserInfo({
    userId
  } = {}) {
    return api.post('/console/v1/getUserInfo', {
      userId
    })
  },
  getToken({
    appKey,
    userId,
    deviceId
  } = {}) {
    return api.get('/v1/getAuthToken', {
      params: {
        appKey,
        userId,
        deviceId
      }
    })
  },listLiveRooms
  ({
    status,
    pageNum,
    pageSize
  } = {}) {
    return api.post('/api/v1/live/list', {
        status,
        pageNum,
        pageSize
    })
  },
  createLiveRoom({
    title,
    notice,
    coverUrl,
    anchorId,
    anchorNick,
    userId,
    EnableLinkMic,
    extension
  } = {}) {
    return api.post('/v1/createLiveRoom', {
      title,
      notice,
      coverUrl,
      anchorId,
      anchorNick,
      userId,
      EnableLinkMic,
      extension
    })
  },
  getLiveRoom({
    id
  } = {}) {
    return api.post('/api/v1/live/get', {
        id
    })
  },
  updateLiveRoom({
    title,
    notice,
    coverUrl,
    anchorId,
    anchorNick,
    userId,
    extension,
    liveId
  } = {}) {
    return api.post('/v1/updateLiveRoom', {
      title,
      notice,
      coverUrl,
      anchorId,
      anchorNick,
      userId,
      extension,
      liveId
    })
  },
  stopLiveRoom({
    liveId,
    userId
  } = {}) {
    return api.post('/v1/stopLiveRoom', {
      liveId,
      userId
    })
  },
  publishLiveRoom({
    liveId,
    userId
  } = {}) {
    return api.post('/v1/publishLiveRoom', {
      liveId,
      userId
    })
  },
  getStandardRoomJumpUrl({
    bizId,
    userId,
    userNick,
    platform
  } = {}) {
    return api.get('/v1/getStandardRoomJumpUrl', {
      params: {
        bizId,
        userId,
        userNick,
        platform: platform || 'web',
        bizType: 'live'
      }
    })
  },
  startLocalPublish({
    fileName,
    pushUrl
  } = {}) {
    return api.post('/v1/startLocalPublish', {
      fileName,
      pushUrl
    })
  },
  getShareUrl({
    liveId
  } = {}) {
    return api.get('/api/getShareUrl', {
      params: {
        liveId
      }
    })
  },
  kickUser({
    appId,
    roomId,
    userId,
    kickUser,
    blockTime
  } = {}) {
    return api.post('/v1/kickRoomUser', {
      appId,
      roomId,
      userId,
      kickUser,
      blockTime
    })
  },
  updatePassword({
    accountId,
    password
  } = {}) {
    return api.post('/users/updatePassword', {
      accountId,
      password
    })
  },
  getAppConfigQRcode({
    deviceType
  } = {}) {
    return api.post('/api/getAppConfigQRCode', {
      deviceType
    })
  },
  getAnchorUrl({
    id
  } = {}) {
    return api.post('/api/getAnchorUrl', {
      id
    })
  }
}
export const service = {
  CasterApiPopService({
    ActionName,
    AppId,
    Parameter,
  } = {}) {
    return api.post('/v1/CasterApi', {
      ActionName,
      AppId,
      Parameter,
    })
  },
};
