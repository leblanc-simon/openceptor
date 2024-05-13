<?php

namespace App\Command;

use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputArgument;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Console\Style\SymfonyStyle;
use Symfony\Component\Mercure\HubInterface;
use Symfony\Component\Mercure\Update;
use Symfony\Component\Notifier\Notification\Notification;
use Symfony\Component\Notifier\NotifierInterface;

#[AsCommand(
    name: 'app:test-notifier',
    description: 'Add a short description for your command',
)]
class TestNotifierCommand extends Command
{

    public function __construct(
        private readonly NotifierInterface $notifier,
        private readonly HubInterface $hub
    )
    {
        parent::__construct();
    }
    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        //$this->notifier->send(new Notification('test', ['chat/mercure_chatter']));
        $this->hub->publish(new Update('7d05faaa-7b1a-4a1e-a696-e0653f729042', json_encode(['test'])));

        return Command::SUCCESS;
    }
}
