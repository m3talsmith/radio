#include <string>
#include <thread>

#include "absl/flags/flag.h"
#include "absl/flags/parse.h"

#include <grpcpp/create_channel.h>
#include <grpcpp/security/credentials.h>

#include "include/StationClient.h"

ABSL_FLAG(std::string, host, "localhost:3001", "GRPC host Address");
ABSL_FLAG(std::string, id, "unknown", "Client ID");

int main(const int argc, char* argv[])
{
    absl::ParseCommandLine(argc, argv);
    const std::string host = absl::GetFlag(FLAGS_host);
    const std::string id = absl::GetFlag(FLAGS_id);

    StationClient client(
        grpc::CreateChannel(host, grpc::InsecureChannelCredentials()),
        id
    );

    while (true)
    {
        std::string message;
        printf("Message (q to quit)? ");
        std::cin >> message;
        if (message == "q")
        {
            exit(0);
        }
        client.say(message);
    }
}