<?php

declare(strict_types=1);

namespace App\Presentation\Step\Controller;

use App\Domain\Step\Dto\UpdateStepPayload;
use App\Domain\Step\Entity\Step;
use App\Domain\Step\Repository\StepRepository;
use App\Domain\Workflow\Repository\WorkflowRepository;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpKernel\Attribute\MapRequestPayload;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;

class UpdateStepController extends AbstractController
{
    public function __construct(
        private readonly StepRepository $stepRepository,
        private readonly WorkflowRepository $workflowRepository,
    ) {
    }

    public function __invoke(#[MapRequestPayload()] UpdateStepPayload $request): JsonResponse
    {
        $workflow = $this->workflowRepository->find($request->workflow);

        if (! $workflow) {
            throw new NotFoundHttpException('Workflow not found');
        }
        foreach ($request->steps as $item) {
            $existingStep = $this->stepRepository->findOneBy(['id' => $item->getId(), 'workflow' => $workflow->getId()]);

            if (! $existingStep) {
                throw new NotFoundHttpException('Step not found');
            }

            $existingStep->setPosition($item->getPosition());
            $this->stepRepository->save($existingStep, true);
        }

        return new JsonResponse(['message' => 'Steps updated successfully']);
    }
}