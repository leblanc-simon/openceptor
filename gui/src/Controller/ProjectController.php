<?php

namespace App\Controller;

use App\Entity\Project;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;
use Symfony\Component\Routing\Requirement\Requirement;

class ProjectController extends AbstractController
{
    #[Route('/{id}', name: 'app_project', requirements: ['id' => Requirement::UUID_V4])]
    public function index(Project $project): Response
    {
        $response = new Response();
        $response->headers->set('Access-Control-Allow-Origin', '*');
        return $this->render('project/index.html.twig', [
            'controller_name' => 'ProjectController',
            'project' => $project,
        ], $response);
    }
}
