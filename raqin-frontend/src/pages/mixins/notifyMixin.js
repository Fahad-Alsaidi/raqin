export default {
    data() {
        return {
            pageTextSaved: {
                message: 'page has been saved successfully',
                color: 'green-4',
                textColor: 'white',
                icon: 'cloud_done'
            },
            networkError: {
                message: 'Server unavailable. Check connectivity.',
                color: 'red-5',
                textColor: 'white',
                icon: 'warning'
            }
        };
    },
    methods:{
        showError: function(msg) {
            return {
                message: msg,
                color: 'red-5',
                textColor: 'white',
                icon: 'warning'
            }
        }
    }
}