package com.lhamacorp.myip.infrastructure;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface IpRepository extends MongoRepository<IpEntity, String> {
}
