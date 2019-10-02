package com.lhamacorp.myip.infrastructure;

import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "ips")
public class IpEntity {

    private String name;
    private String ip;

    public IpEntity(String name, String ip) {
        this.name = name;
        this.ip = ip;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getIp() {
        return ip;
    }

    public void setIp(String ip) {
        this.ip = ip;
    }
}
