package com.lhamacorp.myip.controller;

import com.lhamacorp.myip.infrastructure.IpEntity;
import com.lhamacorp.myip.service.IpService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/ip")
public class IpController {

    private IpService service;

    @Autowired
    public IpController(IpService service) {
        this.service = service;
    }

    @GetMapping
    public ResponseEntity<List<IpEntity>> get() {
        return ResponseEntity.ok(service.get());
    }

    @PostMapping
    public ResponseEntity<String> save(@RequestBody IpEntity entity) {
        service.save(entity);
        return ResponseEntity.status(HttpStatus.CREATED).build();
    }

    @DeleteMapping
    public ResponseEntity<Void> deleteAll() {
        service.delete();
        return ResponseEntity.ok().build();
    }

}
