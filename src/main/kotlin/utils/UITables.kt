package utils

import com.jakewharton.picnic.TextAlignment
import com.jakewharton.picnic.TextBorder
import com.jakewharton.picnic.renderText
import com.jakewharton.picnic.table
import io.kubernetes.client.openapi.models.V1Endpoint
import io.kubernetes.client.openapi.models.V1Endpoints
import io.kubernetes.client.openapi.models.V1Pod
import io.kubernetes.client.openapi.models.V1Service

object UITables {
    val mainMenu = table {
        cellStyle {
            alignment = TextAlignment.MiddleRight
            paddingLeft = 1
            paddingRight = 1
            borderLeft = true
            borderRight = true
        }
        header {
            row {
                cell("Main Menu") {
                    columnSpan = 2
                    alignment = TextAlignment.MiddleCenter
                    border = true
                }
            }
        }
        body {
            row {
                cell("1")
                cell("List Pods")
            }
            row {
                cell("2")
                cell("List Services")
            }
            row {
                cell("3")
                cell("List Endpoints")
                cellStyle {
                    borderBottom = true
                }
            }
        }
        footer {
            row {
                cell("0")
                cell("Exit")
                cellStyle {
                    borderBottom = true
                }
            }
        }
    }.renderText(border = TextBorder.ROUNDED)

    @JvmStatic
    fun podInfoTemplate(data: List<V1Pod>) = table {
        cellStyle {
            alignment = TextAlignment.MiddleRight
            paddingLeft = 1
            paddingRight = 1
            borderLeft = true
            borderRight = true
        }
        header {
            row {
                cell("All Pods") {
                    columnSpan = 8
                    alignment = TextAlignment.MiddleCenter
                    border = true
                }
            }
            row {
                cellStyle {
                    border = true
                    alignment = TextAlignment.BottomLeft
                }
                cell("Index") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Name") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Namespace") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Phase") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Pod IP") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Host IP") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Created") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Started") {
                    alignment = TextAlignment.MiddleCenter
                }
            }
        }
        body {
            data.forEachIndexed { index, it ->
                row {
                    cell(index.toString()) {
                        alignment = TextAlignment.MiddleCenter
                    }
                    cell(it.metadata!!.name) {}
                    cell(it.metadata!!.namespace) {}
                    cell(it.status!!.phase) {}
                    cell(it.status!!.podIP) {}
                    cell(it.status!!.hostIP) {}
                    cell(it.metadata!!.creationTimestamp) {}
                    cell(it.status!!.startTime) {}
                    if (index == data.size - 1) {
                        cellStyle {
                            borderBottom = true
                        }
                    }
                }
            }
        }
    }

    @JvmStatic
    fun serviceInfoTemplate(data: List<V1Service>) = table {
        cellStyle {
            alignment = TextAlignment.MiddleRight
            paddingLeft = 1
            paddingRight = 1
            borderLeft = true
            borderRight = true
        }
        header {
            row {
                cell("All Services") {
                    columnSpan = 6
                    alignment = TextAlignment.MiddleCenter
                    border = true
                }
            }
            row {
                cellStyle {
                    border = true
                    alignment = TextAlignment.BottomLeft
                }
                cell("Index") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Name") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Namespace") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Cluster IP") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Ports") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Created") {
                    alignment = TextAlignment.MiddleCenter
                }
            }
        }
        body {
            data.forEachIndexed { index, it ->
                row {
                    cell(index.toString()) {
                        alignment = TextAlignment.MiddleCenter
                    }
                    cell(it.metadata!!.name) {}
                    cell(it.metadata!!.namespace) {}
                    cell(it.spec!!.clusterIP) {}
                    cell(it.spec!!.ports!!.map { it.port }.joinToString()) {}
                    cell(it.metadata!!.creationTimestamp) {}
                    if (index == data.size - 1) {
                        cellStyle {
                            borderBottom = true
                        }
                    }
                }
            }
        }
    }

    @JvmStatic
    fun endpointInfoTemplate(data: List<V1Endpoints>) = table {
        cellStyle {
            alignment = TextAlignment.MiddleRight
            paddingLeft = 1
            paddingRight = 1
            borderLeft = true
            borderRight = true
        }
        header {
            row {
                cell("All Endpoints") {
                    columnSpan = 6
                    alignment = TextAlignment.MiddleCenter
                    border = true
                }
            }
            row {
                cell("") {
                    columnSpan = 3
                    alignment = TextAlignment.MiddleCenter
                    border = true
                }
                cell("Subsets") {
                    columnSpan = 2
                    alignment = TextAlignment.MiddleCenter
                    border = true
                }
                cell("") {
                    columnSpan = 1
                    alignment = TextAlignment.MiddleCenter
                    border = true
                }
            }
            row {
                cellStyle {
                    border = true
                    alignment = TextAlignment.BottomLeft
                }
                cell("Index") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Name") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Namespace") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Adresses") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Ports") {
                    alignment = TextAlignment.MiddleCenter
                }
                cell("Created") {
                    alignment = TextAlignment.MiddleCenter
                }
            }
        }
        body {
            data.forEachIndexed { index, it ->
                row {
                    cell(index.toString()) {
                        alignment = TextAlignment.MiddleCenter
                    }
                    cell(it.metadata!!.name) {}
                    cell(it.metadata!!.namespace) {}
                    cell(it.subsets!!.flatMap { it -> it.addresses!!.map { it.ip } }.joinToString()) {}
                    cell(it.subsets!!.flatMap { it -> it.ports!!.map { it.port } }.joinToString()) {}
                    cell(it.metadata!!.creationTimestamp) {}
                    if (index == data.size - 1) {
                        cellStyle {
                            borderBottom = true
                        }
                    }
                }
            }
        }
    }
}