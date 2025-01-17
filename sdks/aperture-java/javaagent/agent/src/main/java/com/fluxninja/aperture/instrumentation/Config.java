package com.fluxninja.aperture.instrumentation;

import com.fluxninja.aperture.sdk.ApertureSDK;
import com.fluxninja.aperture.sdk.ApertureSDKBuilder;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.Properties;

public class Config {
    public static final String CONFIG_FILENAME_PROPERTY = "aperture.javaagent.config.file";

    public static final String AGENT_HOST_PROPERTY = "aperture.agent.hostname";
    public static final String AGENT_PORT_PROPERTY = "aperture.agent.port";
    public static final String CONNECTION_TIMEOUT_MILLIS_PROPERTY = "aperture.connection.timeout.millis";
    public static final String BLOCKED_PATHS_PROPERTY = "aperture.javaagent.blocked.paths";
    public static final String BLOCKED_PATHS_REGEX_PROPERTY = "aperture.javaagent.blocked.paths.regex";

    private static final String AGENT_HOST_DEFAULT_VALUE = "localhost";
    private static final String AGENT_PORT_DEFAULT_VALUE = "8089";
    private static final String CONNECTION_TIMEOUT_MILLIS_DEFAULT_VALUE = "1000";
    private static final String BLOCKED_PATHS_DEFAULT_VALUE = "";
    private static final String BLOCKED_PATHS_REGEX_DEFAULT_VALUE = "false";


    static private final List<String> allProperties = new ArrayList<String>() {{
        add(AGENT_HOST_PROPERTY);
        add(AGENT_PORT_PROPERTY);
        add(CONNECTION_TIMEOUT_MILLIS_PROPERTY);
        add(BLOCKED_PATHS_PROPERTY);
        add(BLOCKED_PATHS_REGEX_PROPERTY);
    }};

    static Properties loadProperties() {
        Properties props = new Properties();
        String configFileName = System.getProperty(CONFIG_FILENAME_PROPERTY);
        if (configFileName == null) {
            configFileName = System.getenv(envNameFromPropertyName(CONFIG_FILENAME_PROPERTY));
        }
        try {
            if (configFileName != null) {
                props.load(Files.newInputStream(Paths.get(configFileName)));
            }
        } catch (IOException e) {
            throw new RuntimeException("Could not read properties from file", e);
        }

        // Get property overrides from env and commandline
        for (String key: allProperties) {
            String val = getFromEnv(key);
            if (val != null) {
                props.put(key, val);
            }
        }

        return props;
    }

    static String getFromEnv(String name) {
        // Read system property; If not set, use env variable.
        String systemProperty = System.getProperty(name);
        if (systemProperty != null) {
            return systemProperty;
        }
        String envVariableName = envNameFromPropertyName(name);
        return System.getenv(envVariableName);
    }

    public static ApertureSDK newSDKFromConfig() {
        ApertureSDKBuilder builder = ApertureSDK.builder();
        Properties config = loadProperties();
        ApertureSDK sdk;
        try {
            sdk = builder
                    .setHost(config.getProperty(AGENT_HOST_PROPERTY, AGENT_HOST_DEFAULT_VALUE))
                    .setPort(Integer.parseInt(config.getProperty(AGENT_PORT_PROPERTY, AGENT_PORT_DEFAULT_VALUE)))
                    .setDuration(Duration.ofMillis(Integer.parseInt(config.getProperty(CONNECTION_TIMEOUT_MILLIS_PROPERTY, CONNECTION_TIMEOUT_MILLIS_DEFAULT_VALUE))))
                    .addBlockedPaths(config.getProperty(BLOCKED_PATHS_PROPERTY, BLOCKED_PATHS_DEFAULT_VALUE))
                    .setBlockedPathMatchRegex(Boolean.parseBoolean(config.getProperty(BLOCKED_PATHS_REGEX_PROPERTY, BLOCKED_PATHS_REGEX_DEFAULT_VALUE)))
                    .build();
        } catch (Exception e) {
            throw new RuntimeException("failed to create Aperture SDK from config", e);
        }
        return sdk;
    }

    private static String envNameFromPropertyName(String propertyName) {
        return propertyName.toUpperCase().replace(".", "_");
    }
}
