FROM openjdk:8-jre

COPY . /flyway
WORKDIR /flyway

ENV FLYWAY_VERSION 5.0.7

RUN curl -L https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/${FLYWAY_VERSION}/flyway-commandline-${FLYWAY_VERSION}.tar.gz -o flyway-commandline-${FLYWAY_VERSION}.tar.gz \
  && tar -xzf flyway-commandline-${FLYWAY_VERSION}.tar.gz --strip-components=1 \
  && rm flyway-commandline-${FLYWAY_VERSION}.tar.gz \
  && ln -s /flyway/flyway /usr/local/bin/flyway

RUN ["chmod", "+x", "wait-for-it.sh" ]
CMD [ "./wait-for-it.sh", "db:5432" ]

ENTRYPOINT [ "flyway" ]
