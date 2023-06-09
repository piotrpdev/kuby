import io.kubernetes.client.openapi.ApiException
import io.kubernetes.client.openapi.Configuration
import io.kubernetes.client.openapi.apis.CoreV1Api
import io.kubernetes.client.util.Config
import java.io.IOException

fun main(args: Array<String>) {
    listPods()
}

fun listPods() {
    val client = Config.defaultClient()
    Configuration.setDefaultApiClient(client)
    val api = CoreV1Api()
    val list = api.listPodForAllNamespaces(null, null, null, null, null, null, null, null, null, false)
    for (item in list.items) {
        println(item.metadata!!.name)
    }
}