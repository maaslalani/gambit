FROM scratch
COPY gambit /usr/local/bin/gambit

# Expose data volume
VOLUME /data

# Expose ports
EXPOSE 53531/tcp

# Set the default command
ENTRYPOINT [ "/usr/local/bin/gambit" ]