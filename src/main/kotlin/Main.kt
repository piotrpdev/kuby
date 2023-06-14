import com.jakewharton.picnic.TextBorder
import com.jakewharton.picnic.renderText
import com.varabyte.kotter.foundation.anim.text
import com.varabyte.kotter.foundation.anim.textAnimOf
import com.varabyte.kotter.foundation.anim.textLine
import com.varabyte.kotter.foundation.input.input
import com.varabyte.kotter.foundation.input.onInputChanged
import com.varabyte.kotter.foundation.input.onInputEntered
import com.varabyte.kotter.foundation.input.runUntilInputEntered
import com.varabyte.kotter.foundation.liveVarOf
import com.varabyte.kotter.foundation.session
import com.varabyte.kotter.foundation.text.text
import com.varabyte.kotter.foundation.text.textLine
import io.kubernetes.client.openapi.ApiClient
import io.kubernetes.client.openapi.ApiException
import io.kubernetes.client.openapi.Configuration
import io.kubernetes.client.openapi.apis.CoreV1Api
import io.kubernetes.client.util.Config
import mu.KotlinLogging
import utils.KubeClient
import utils.UITables
import java.io.IOException
import kotlin.system.exitProcess
import kotlin.time.Duration.Companion.milliseconds

val logger = KotlinLogging.logger {}

fun listPods(): String {
    logger.debug { "listPods() function invoked" }
    return try {
        val list = KubeClient.getApi().listPodForAllNamespaces(null, null, null, null, null, null, null, null, null, false)
        UITables.podInfoTemplate(list.items).renderText(border = TextBorder.ROUNDED)
    } catch (e: Exception) {
        logger.error(e) { "Error invoking listPodForAllNamespaces()" }
        "Error! Have you started docker? ${e.message}"
    }
}

fun listServices(): String {
    logger.debug { "listServices() function invoked" }
    return try {
        val list = KubeClient.getApi().listServiceForAllNamespaces(null, null, null, null, null, null, null, null, null, null)
        UITables.serviceInfoTemplate(list.items).renderText(border = TextBorder.ROUNDED)
    } catch (e: Exception) {
        logger.error(e) { "Error invoking listServiceForAllNamespaces()" }
        "Error! Have you started docker? ${e.message}"
    }
}

fun listEndpoints(): String {
    logger.debug { "listEndpoints() function invoked" }
    return try {
        val list = KubeClient.getApi().listEndpointsForAllNamespaces(null, null, null, null, null, null, null, null, null, null)
        UITables.endpointInfoTemplate(list.items).renderText(border = TextBorder.ROUNDED)
    } catch (e: Exception) {
        logger.error(e) { "Error invoking listEndpointsForAllNamespaces()" }
        "Error! Have you started docker? ${e.message}"
    }
}

fun exitApp(): String {
    logger.debug { "exitApp() function invoked" }

    logger.debug { "Exiting...bye" }
    return "Exiting...bye"
    //exitProcess(0)
}

fun main() {
    logger.debug { "main() function invoked" }

    session {
        var firstApiCallMade = false
        val spinnerAnim = textAnimOf(listOf("\\", "|", "/", "-"), 125.milliseconds)

        section {
            textLine(
                """
                ██╗  ██╗██╗   ██╗██████╗ ██╗   ██╗
                ██║ ██╔╝██║   ██║██╔══██╗╚██╗ ██╔╝
                █████╔╝ ██║   ██║██████╔╝ ╚████╔╝ 
                ██╔═██╗ ██║   ██║██╔══██╗  ╚██╔╝  
                ██║  ██╗╚██████╔╝██████╔╝   ██║   
                ╚═╝  ╚═╝ ╚═════╝ ╚═════╝    ╚═╝   
                
                """.trimIndent()
            )
            textLine(UITables.mainMenu)
        }.run()

        do {
            var choice: Int? = null
            var choiceResult by liveVarOf<String>("")

            section {
                textLine()
                text("Enter option: "); input()
            }.runUntilInputEntered {
                onInputChanged { input = input.filter { it.isDigit() } }
                onInputEntered { choice = input.toIntOrNull() }
            }

            section {
                val stillCalculating = choiceResult.isEmpty()

                // TODO: Bit of a hack, find a better way to know when the first API call has been made
                if (!firstApiCallMade && choice in 1..3) {
                    if (stillCalculating) {
                        text(spinnerAnim)
                    } else {
                        text("✓")
                        firstApiCallMade = true
                    }
                    textLine(" Connection to Kubernetes API")
                }

                if (!stillCalculating) { textLine(choiceResult) }
            }.run {
                choiceResult = when (choice) {
                    1 -> listPods()
                    2 -> listServices()
                    3 -> listEndpoints()
                    0 -> exitApp()
                    else -> "Invalid option entered: $choice"
                }
            }
        } while (choice != 0)
    }
}