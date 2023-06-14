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
    implementation("com.varabyte.kotter:kotter-jvm:1.1.0")
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

tasks.jar {
    manifest.attributes["Main-Class"] = "MainKt"
    // for building a fat jar - include all dependencies
    duplicatesStrategy = DuplicatesStrategy.EXCLUDE
    from(sourceSets.main.get().output)
    dependsOn(configurations.runtimeClasspath)
    from({
        configurations.runtimeClasspath.get().filter { it.name.endsWith("jar") }.map { zipTree(it) }
    }) {
        // https://stackoverflow.com/questions/999489/invalid-signature-file-when-attempting-to-run-a-jar
        exclude("META-INF/*.SF")
        exclude("META-INF/*.DSA")
        exclude("META-INF/*.RSA")
    }
}