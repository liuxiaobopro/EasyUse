import { ElNotification, ElMessageBox } from 'element-plus'

export function elNotification(msg, type = 'success') {
    return ElNotification({
        title: type,
        message: msg || '未知错误',
        type: type,
        duration: 3000,
    })
}

export function elMessageBox(msg, confirmCallback, cancelCallback = () => { }) {
    return ElMessageBox.confirm(
        msg,
        'Warning',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
            confirmButtonClass: 'ExitConfirmButton',
            cancelButtonClass: 'ExitCancelButton',
            center: true,
        }
    )
        .then(confirmCallback)
        .catch(cancelCallback)
}