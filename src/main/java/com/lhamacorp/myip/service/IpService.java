package com.lhamacorp.myip.service;

import com.lhamacorp.myip.infrastructure.IpEntity;
import com.lhamacorp.myip.infrastructure.IpRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IpService {

    private IpRepository repository;

    @Autowired
    public IpService(IpRepository repository) {
        this.repository = repository;
    }

    public List<IpEntity> get() {
        return repository.findAll();
    }

    public void save(IpEntity entity) {
        repository.save(entity);
    }

    public void delete() {
        repository.deleteAll();
    }

}
