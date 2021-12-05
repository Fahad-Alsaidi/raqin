export default {
    methods: {
        showError: function (msg) {
            return {
                message: msg,
                color: 'red-5',
                textColor: 'white',
                icon: 'warning'
            }
        },
        showSuccess: function (msg) {
            return {
                type: 'positive',
                color: 'positive',
                textColor: 'white',
                icon: 'cloud_done',
                timeout: 3000,
                position: 'bottom',
                message: msg
            }
        }
    }
}