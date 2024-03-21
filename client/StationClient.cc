#include <chrono>
#include <iostream>
#include <memory>
#include <string>
#include <thread>
#include <utility>

#include <grpc/grpc.h>
#include <grpcpp/channel.h>
#include <grpcpp/client_context.h>

#include "radio.grpc.pb.h"
#include "include/StationClient.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientReaderWriter;
using grpc::Status;

using radio::RadioAPI;
using radio::Request;
using radio::Broadcast;

StationClient::StationClient(
    const std::shared_ptr<Channel>& channel, std::string  caller_id) :
        caller_id(std::move(caller_id)),
        stub_(RadioAPI::NewStub(channel))
{
    printf("turning station on...");

    on = true;
    connection = stub_->Station(&context);

    station_thread = std::thread([this]
    {
        say("joining station");
        while (on)
        {
            Broadcast reply;
            if (!connection->Read(&reply))
            {
                printf("error reading connection\n");
                off();
                break;
            }
            printf("[INCOMING] %s\r\n", reply.message().data());
            std::this_thread::sleep_for(
                std::chrono::milliseconds(100)
            );
        }
    });

    printf("done\n");
}

StationClient::~StationClient()
{
    off();
}

void StationClient::off()
{
    printf("turning station off...");
    on = false;
    if (station_thread.joinable()) station_thread.join();
    connection->WritesDone();
    if (!connection->Finish().ok())
    {
        printf("error closing...");
    }
    printf("done\n");
}

void StationClient::say(const std::string& message)
{
    Request request;
    request.set_caller_id(caller_id);
    request.set_message(message);

    if (!connection->Write(request))
    {
        printf("error sending message: %s\n", message.data());
    }
}



