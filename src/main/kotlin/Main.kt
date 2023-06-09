import io.kubernetes.client.openapi.ApiClient
import io.kubernetes.client.openapi.ApiException
import io.kubernetes.client.openapi.Configuration
import io.kubernetes.client.openapi.apis.CoreV1Api
import io.kubernetes.client.util.Config
import mu.KotlinLogging
import utils.UITables
import java.io.IOException
import kotlin.system.exitProcess

private val logger = KotlinLogging.logger {}

object KubeClient {
    private val client: ApiClient = Config.defaultClient()
    private var api: CoreV1Api

    init {
        Configuration.setDefaultApiClient(client)
        api = CoreV1Api(client)
    }

    fun getApi(): CoreV1Api {
        return api
    }
}

fun listPods() {
    logger.debug { "listPods() function invoked" }
    val list = KubeClient.getApi().listPodForAllNamespaces(null, null, null, null, null, null, null, null, null, false)
    println(UITables.podInfoTemplate(list.items))
}

fun listServices() {
    logger.debug { "listServices() function invoked" }
    val list = KubeClient.getApi().listServiceForAllNamespaces(null, null, null, null, null, null, null, null, null, null)
    println(UITables.serviceInfoTemplate(list.items))
}

fun listEndpoints() {
    logger.debug { "listEndpoints() function invoked" }
    val list = KubeClient.getApi().listEndpointsForAllNamespaces(null, null, null, null, null, null, null, null, null, null)
    println(UITables.endpointInfoTemplate(list.items))
}

fun mainMenu(): Int? {
    logger.debug { "mainMenu() function invoked" }

    println(UITables.mainMenu)

    print("Enter option: ")

    return readln().toIntOrNull()
}

fun exitApp() {
    logger.debug { "exitApp() function invoked" }

    logger.debug { "Exiting...bye" }
    exitProcess(0)
}

fun runMenu() {
    logger.debug { "runMenu() function invoked" }

    do {
        when (val option = mainMenu()) {
            1 -> listPods()
            2 -> listServices()
            3 -> listEndpoints()
            0 -> exitApp()
            else -> println("Invalid option entered: $option")
        }
    } while (true)
}

fun main() {
    logger.debug { "main() function invoked" }
    // https://patorjk.com/software/taag/
    println(
        """
            ██╗  ██╗██╗   ██╗██████╗ ██╗   ██╗
            ██║ ██╔╝██║   ██║██╔══██╗╚██╗ ██╔╝
            █████╔╝ ██║   ██║██████╔╝ ╚████╔╝ 
            ██╔═██╗ ██║   ██║██╔══██╗  ╚██╔╝  
            ██║  ██╗╚██████╔╝██████╔╝   ██║   
            ╚═╝  ╚═╝ ╚═════╝ ╚═════╝    ╚═╝   
            
        """.trimIndent()
    )

    runMenu()
}