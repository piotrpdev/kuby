plugins {
    kotlin("jvm") version "1.8.21"
    application
}

group = "com.redhat"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation("io.github.microutils:kotlin-logging:3.0.5")
    implementation("org.slf4j:slf4j-simple:2.0.5")
    testImplementation(kotlin("test"))
    implementation("io.kubernetes:client-java:18.0.0")
    implementation("com.jakewharton.picnic:picnic:0.6.0")
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    jvmToolchain(11)
}

application {
    mainClass.set("MainKt")
}
