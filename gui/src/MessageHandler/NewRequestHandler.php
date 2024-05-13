<?php

namespace App\MessageHandler;

use App\Domain\Request\Request;
use Symfony\Component\Mercure\HubInterface;
use Symfony\Component\Mercure\Update;
use Symfony\Component\Messenger\Attribute\AsMessageHandler;
use Symfony\Component\Serializer\Serializer;
use Symfony\Component\Serializer\SerializerInterface;

#[AsMessageHandler]
class NewRequestHandler
{
    public function __construct(
        private readonly SerializerInterface $serializer,
        private readonly HubInterface $hub
    ) {
    }

    public function __invoke(Request $request)
    {
        dump($request);
        $this->hub->publish(new Update($request->getProjectId(), $this->serializer->serialize($request, 'json')));
    }
}
