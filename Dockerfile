FROM quay.io/minio/minio
LABEL authors="design"
WORKDIR /apphome
COPY . .
RUN mkdir /storage
VOLUME /storage
ENV MINIO_ROOT_USER=admin
ENV MINIO_ROOT_PASWORD=admin12345
ENV MINIO_VOLUMES="/sotage"
EXPOSE 9000
EXPOSE 9090

CMD ["minio","server","/storage","--console-address",":9090"]