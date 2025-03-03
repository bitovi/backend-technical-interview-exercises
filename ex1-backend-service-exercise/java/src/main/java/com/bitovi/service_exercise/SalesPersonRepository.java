package com.bitovi.service_exercise;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.io.Resource;
import org.springframework.stereotype.Repository;

import java.io.IOException;
import java.util.List;
import java.util.Objects;

@Repository
public class SalesPersonRepository {
    static List<SalesPerson> records = null;

    public SalesPersonRepository(
            @Value("classpath:records.json")
            Resource recordsFile,
            @Autowired
            ObjectMapper objectMapper
    ) {
        try {
            records = objectMapper.createParser(recordsFile.getInputStream())
                    .readValueAs(new TypeReference<List<SalesPerson>>(){});
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public List<SalesPerson> findAll() {
        return records;
    }

    public record SalesPerson(
            String name,
            List<SalesPersonSalesAggregate> sales
    ) {
        public SalesPerson {
            Objects.requireNonNull(name);
            Objects.requireNonNull(sales);
        }
    }

    public record SalesPersonSalesAggregate(
            Integer month,
            Integer amount
    ) {
        public SalesPersonSalesAggregate {
            Objects.requireNonNull(month);
            Objects.requireNonNull(amount);
        }
    }
}