Notes:
1. We created an exporter using https://github.com/Chinwendu20/otel_components_generator
2. Installed OCB tool from https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder
3. Added a logic to send logs to different endpoints based on log's resource attributes

To configure custom collector
```
  builder  --config=otelcol-builder.yaml
```

To run the collector
```
/tmp/dist/otelcol-custom --config=otel-config.yaml
```


1. We looked up at https://github.com/open-telemetry/opentelemetry-proto to understand the expected request for GRPC calls
  - For making a GRPC api call through postman
    - Download the repository https://github.com/open-telemetry/opentelemetry-proto
    - load the proto file opentelemetry/proto/collector/logs/v1/logs_service.proto
    - Give path to the project ${pwd}
    - use example mesage 
    ```
    {
    "resource_logs": [
        {
            "schema_url": "consequat dolore qui ut",
            "scope_logs": [
                {
                    "log_records": [
                        {
                            "attributes": [
                                {
                                    "value": {
                                        "int_value": "1681894",
                                        "kvlist_value": {
                                            "values": [
                                                {
                                                    "value": {
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "string_value": "ad pariatur labore eiusmod",
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "double_value": "Infinity",
                                                        "bool_value": true,
                                                        "bytes_value": "fAYLqveV/lxdJ7t9FS4L",
                                                        "int_value": "3"
                                                    },
                                                    "key": "ut"
                                                },
                                                {
                                                    "value": {
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "double_value": -72077666.36282252,
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "bytes_value": "2qsvHv99TdfKHzJoTmD4",
                                                        "string_value": "esse nisi do",
                                                        "bool_value": true,
                                                        "int_value": "3"
                                                    },
                                                    "key": "nulla ex consequat"
                                                },
                                                {
                                                    "key": "anim",
                                                    "value": {
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "bytes_value": "b0Ueg/2DOsS+YU/HsLv9ozlYPKZV+XtJVEFN2VWC",
                                                        "bool_value": true,
                                                        "string_value": "ad",
                                                        "double_value": "Infinity",
                                                        "int_value": "643446"
                                                    }
                                                },
                                                {
                                                    "key": "adipisicing",
                                                    "value": {
                                                        "string_value": "deserunt enim nisi commodo",
                                                        "double_value": 38521165.98134762,
                                                        "bytes_value": "L+",
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "int_value": "36",
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "bool_value": true
                                                    }
                                                },
                                                {
                                                    "key": "sint commodo nisi",
                                                    "value": {
                                                        "string_value": "ex cillum ipsum adipisicing",
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "double_value": 46419268.95817354,
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "bytes_value": "Z7uYcXS5L5DpneaNtUWLQ+t8QwzzDZt/4rWchTPfe/t",
                                                        "int_value": "35751259732",
                                                        "bool_value": false
                                                    }
                                                }
                                            ]
                                        },
                                        "string_value": "Duis",
                                        "array_value": {
                                            "values": []
                                        },
                                        "bytes_value": "Tt+V",
                                        "bool_value": false,
                                        "double_value": 92846087.08895174
                                    },
                                    "key": "est qui"
                                },
                                {
                                    "value": {
                                        "bool_value": true,
                                        "int_value": "67114496",
                                        "string_value": "ipsum Ut",
                                        "bytes_value": "P/T5b7Z2ggVGdVyzv51mupycTkUZasJ9FhIq",
                                        "double_value": "NaN",
                                        "kvlist_value": {
                                            "values": [
                                                {
                                                    "key": "ad",
                                                    "value": {
                                                        "int_value": "5766415750",
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "string_value": "reprehenderit dolore",
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "bytes_value": "wfxA7YGHtuM/Xeg",
                                                        "double_value": "-Infinity",
                                                        "bool_value": true
                                                    }
                                                },
                                                {
                                                    "value": {
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "bool_value": false,
                                                        "int_value": "6020",
                                                        "bytes_value": "IyeiEG+yMirDdSHgcjv",
                                                        "string_value": "aute qui et deserunt cillum",
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "double_value": "Infinity"
                                                    },
                                                    "key": "esse non laborum Lorem"
                                                }
                                            ]
                                        },
                                        "array_value": {
                                            "values": []
                                        }
                                    },
                                    "key": "magna"
                                },
                                {
                                    "value": {
                                        "bool_value": false,
                                        "bytes_value": "XaEd9/Ce0+w2K8MCavBUyNWenPR/QFZx/580FsCY",
                                        "int_value": "86",
                                        "kvlist_value": {
                                            "values": [
                                                {
                                                    "value": {
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "int_value": "177607",
                                                        "bool_value": false,
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "string_value": "Lorem non tempor",
                                                        "double_value": 5622706.076719061,
                                                        "bytes_value": "e5Vb"
                                                    },
                                                    "key": "proident magna dolor eu"
                                                },
                                                {
                                                    "value": {
                                                        "string_value": "dolor magna",
                                                        "bytes_value": "KTRdVggrigGL1+PN7U7HiADdXZh7XweJKwagMGCKk8l=",
                                                        "bool_value": false,
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "int_value": "30492",
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "double_value": "Infinity"
                                                    },
                                                    "key": "laborum in"
                                                },
                                                {
                                                    "key": "mollit in enim",
                                                    "value": {
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "int_value": "3825338634",
                                                        "string_value": "ad",
                                                        "double_value": "NaN",
                                                        "bool_value": true,
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "bytes_value": "Bsx9SI6yvZJ6CS=="
                                                    }
                                                },
                                                {
                                                    "value": {
                                                        "double_value": "Infinity",
                                                        "bytes_value": "NdG+rM5A",
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "string_value": "ipsum nulla",
                                                        "bool_value": false,
                                                        "int_value": "997362",
                                                        "array_value": {
                                                            "values": []
                                                        }
                                                    },
                                                    "key": "officia Lorem pariatur incididunt"
                                                }
                                            ]
                                        },
                                        "array_value": {
                                            "values": []
                                        },
                                        "double_value": "-Infinity",
                                        "string_value": "eiusmod nostrud"
                                    },
                                    "key": "sunt sed ad nostrud"
                                },
                                {
                                    "value": {
                                        "kvlist_value": {
                                            "values": [
                                                {
                                                    "key": "ea consequat magna ut dolor",
                                                    "value": {
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "int_value": "928120620",
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "string_value": "incididunt",
                                                        "double_value": 52270510.90201238,
                                                        "bytes_value": "scdo3Ma3k7iOGpHN1FouF10MgueXfR1/BDa=",
                                                        "bool_value": true
                                                    }
                                                },
                                                {
                                                    "key": "et laboris",
                                                    "value": {
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "double_value": "NaN",
                                                        "string_value": "dolore ut ex sint",
                                                        "int_value": "65171",
                                                        "bool_value": false,
                                                        "bytes_value": "KAEvDnIrXO=="
                                                    }
                                                },
                                                {
                                                    "value": {
                                                        "bool_value": false,
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "double_value": -6981963.266975284,
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "bytes_value": "FSS8E3kQoWXP",
                                                        "string_value": "culpa laborum in",
                                                        "int_value": "726"
                                                    },
                                                    "key": "elit id aute in sunt"
                                                }
                                            ]
                                        },
                                        "string_value": "aute consectetur deserunt",
                                        "double_value": 9559796.854001418,
                                        "bytes_value": "3jgHqDqQOlgC4+9fkfAy",
                                        "bool_value": false,
                                        "int_value": "172",
                                        "array_value": {
                                            "values": []
                                        }
                                    },
                                    "key": "laborum nostrud"
                                },
                                {
                                    "key": "in",
                                    "value": {
                                        "bytes_value": "2Ez6OUoF0Lm9HADuZWNQm+Z/dHhUwydcpT",
                                        "string_value": "Excepteur in consequat sunt",
                                        "bool_value": true,
                                        "int_value": "285320",
                                        "kvlist_value": {
                                            "values": [
                                                {
                                                    "key": "voluptate velit amet consectetur laborum",
                                                    "value": {
                                                        "double_value": "-Infinity",
                                                        "int_value": "99",
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "string_value": "do amet dolor eiusmod",
                                                        "bool_value": true,
                                                        "kvlist_value": {
                                                            "values": []
                                                        },
                                                        "bytes_value": "8uA9v1jehgRdQLgrGtSdO8tRxWjrMBnw"
                                                    }
                                                },
                                                {
                                                    "key": "nostrud pariatur eiusmod amet incididunt",
                                                    "value": {
                                                        "int_value": "678",
                                                        "bool_value": true,
                                                        "array_value": {
                                                            "values": []
                                                        },
                                                        "double_value": -61500830.895585336,
                                                        "string_value": "Ut fugiat ad voluptate laboris",
                                                        "bytes_value": "yNmnF973",
                                                        "kvlist_value": {
                                                            "values": []
                                                        }
                                                    }
                                                }
                                            ]
                                        },
                                        "array_value": {
                                            "values": []
                                        },
                                        "double_value": "-Infinity"
                                    }
                                }
                            ],
                            "body": {
                                "double_value": -9508829.242622554,
                                "bool_value": true,
                                "bytes_value": "tymbxxeJ",
                                "string_value": "Lorem consequat magna est occaecat",
                                "array_value": {
                                    "values": []
                                },
                                "int_value": "59059"
                            },
                            "observed_time_unix_nano": "09296",
                            "dropped_attributes_count": 2404104320,
                            "severity_number": 9,
                            "time_unix_nano": "1197147907",
                            "flags": 1740215907,
                            "span_id": [72, 101, 108, 108, 111, 32, 87, 111],
                            "trace_id": [72, 101, 108, 108, 111, 32, 87, 111, 114, 108,100, 0, 0, 0, 0, 0],
                            "severity_text": "est" 
                            }
                    ],
                    "schema_url": "in culpa commodo",
                    "scope": {
                        "version": "reprehenderit",
                        "dropped_attributes_count": 2231983473,
                        "attributes": [],
                        "name": "aute tempor mollit ea deserunt"
                    }
                }
            ],
            "resource": {
                "attributes": [
                    {
                        "value": {
                            "string_value": "dummyProjectU"
                        },
                        "key": "project_id"
                    }
                ],
                "dropped_attributes_count": 4284046366
            }
        }
        
    ]
} 
    ```

Expectation
1. We are getting the project_id from the resource attribute
2. Looking up the endpoint data from Redis using the project_id identifier 
3. Send the logs to different endpoint