//Prompt model
function Prompt() {
    //toast is toast from SweetAlert2
    let toast = function (c) {
        //parameter
        const {
            msg = "",
            icon = "success",
            position = "top-end",
        } = c

        //Toast from SweetAlert2
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            icon: icon,
            position: position,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({})
    }
    //success is the modal from SweetAlert2 with icon Success
    let success = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c
        Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer: footer
        })
    }
    //error is the modal from SweerAlert2 with icon Error
    let error = function (c) {
        const {
            msg = "",
            title = "",
            footer = "",
        } = c
        Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer: footer,
        })
    }

    async function custom(c) {
        const {
            icon = "",
            msg = "",
            title = "",
            showConfirmButton = true,
        } = c;
        const { value: result } = await Swal.fire({
            showConfirmButton: showConfirmButton,
            icon: icon,
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            customClass: {
                htmlContainer: 'overflow-visible'
            },
            willOpen: () => {
                if (c.willOpen !== undefined){
                    c.willOpen();
                }
            },
            didOpen: () => {
                document.getElementById('start').disabled = false;
                document.getElementById('end').disabled = false;
            }
            ,
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            }
        })
        if (result) {
            //if user click ok button
            if (result.dismiss !== Swal.DismissReason.cancel ) {
                //if form has value
                if (result.value !== "") {
                    //if form has a callback function
                    if (c.callback !== undefined){
                        //execute the callback function
                        c.callback(result);
                    }
                }else {
                    c.callback(false)
                }
            }else {
                c.callback(false)
            }
        }

    }
    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}