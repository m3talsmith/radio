#ifndef STATION_H
#define STATION_H


#include <grpc/grpc.h>
#include <grpcpp/channel.h>

#include "../radio.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReaderWriter;

using radio::RadioAPI;
using radio::Request;
using radio::Broadcast;

class StationClient {
    public:
        std::string caller_id;
        bool on;

        ClientContext context;

        std::unique_ptr<ClientReaderWriter<Request,Broadcast>> connection;
        std::thread station_thread;

        StationClient(
            const std::shared_ptr<Channel>& channel,
            std::string caller_id
        );
        ~StationClient();

        void off();
        void say(const std::string& message);
    private:
        std::unique_ptr<RadioAPI::Stub> stub_;
};

#endif // STATION_H
