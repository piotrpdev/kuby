package utils

import io.kubernetes.client.openapi.ApiClient
import io.kubernetes.client.openapi.Configuration
import io.kubernetes.client.openapi.apis.CoreV1Api
import io.kubernetes.client.util.Config
import logger

object KubeClient {
    private val client: ApiClient = Config.defaultClient()
    private var api: CoreV1Api

    init {
        logger.debug { "KubeClient init block" }
        Configuration.setDefaultApiClient(client)
        logger.debug { "Creating KubeClient API instance" }
        api = CoreV1Api(client)
    }

    fun getApi(): CoreV1Api {
        return api
    }
}