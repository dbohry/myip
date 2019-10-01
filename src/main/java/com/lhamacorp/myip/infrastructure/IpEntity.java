package com.lhamacorp.myip.infrastructure;

import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "ips")
public class IpEntity {

    public IpEntity() {
    }

    public IpEntity(String ip) {
        this.ip = ip;
    }

    private String ip;

    public String getIp() {
        return ip;
    }

    public void setIp(String ip) {
        this.ip = ip;
    }
}
