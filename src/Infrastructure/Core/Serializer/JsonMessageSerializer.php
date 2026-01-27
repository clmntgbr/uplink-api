<?php

declare(strict_types=1);

namespace App\Infrastructure\Core\Serializer;

use App\Shared\Application\Command\AsynchronousCoreInterface;
use Override;
use RuntimeException;
use Symfony\Component\DependencyInjection\Attribute\Autowire;
use Symfony\Component\Messenger\Envelope;
use Symfony\Component\Messenger\Transport\Serialization\SerializerInterface;

use function Safe\json_encode;

class JsonMessageSerializer implements SerializerInterface
{
    public function __construct(
        #[Autowire('%env(WEBHOOK_URI)%')]
        private readonly string $webhookUri,
    ) {
    }

    /**
     * @param array<string, mixed> $encodedEnvelope
     */
    #[Override]
    public function decode(array $encodedEnvelope): Envelope
    {
        throw new RuntimeException('Decode not implemented');
    }

    /**
     * @return array<string, mixed>
     */
    #[Override]
    public function encode(Envelope $envelope): array
    {
        $message = $envelope->getMessage();

        if (! $message instanceof AsynchronousCoreInterface) {
            throw new RuntimeException('The message must implement AsynchronousCoreInterface.');
        }

        $data = [
            'class' => $message::class,
            'payload' => $message->jsonSerialize(),
            'webhook_url_success' => $this->webhookUri . '/' . $message->getWebhookUrlSuccess(),
            'webhook_url_failure' => $this->webhookUri . '/' . $message->getWebhookUrlFailure(),
        ];

        return [
            'body' => json_encode($data),
            'headers' => [
                'Content-Type' => 'application/json',
            ],
        ];
    }
}
