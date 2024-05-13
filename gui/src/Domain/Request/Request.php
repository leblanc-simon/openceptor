<?php

namespace App\Domain\Request;

use Symfony\Component\Serializer\Annotation\Ignore;

class Request
{
    private ?string $projectId;
    private ?string $httpVersion;
    private ?string $host;
    private ?string $method;
    private ?string $requestUri;
    private ?array $headers;
    private Cookies $cookies;
    private ?string $queryString;
    private ?string $postValues;
    private ?array $formValues;
    private ?array $formFiles;
    private ?string $remoteAddr;


    public function getProjectId(): ?string
    {
        return $this->projectId;
    }

    /**
     * @param string|null $projectId
     */
    public function setProjectId(?string $projectId): void
    {
        $this->projectId = $projectId;
    }

    /**
     * @return string|null
     */
    public function getHttpVersion(): ?string
    {
        return $this->httpVersion;
    }

    /**
     * @param string|null $httpVersion
     */
    public function setHttpVersion(?string $httpVersion): void
    {
        $this->httpVersion = $httpVersion;
    }

    /**
     * @return string|null
     */
    public function getHost(): ?string
    {
        return $this->host;
    }

    /**
     * @param string|null $host
     */
    public function setHost(?string $host): void
    {
        $this->host = $host;
    }

    /**
     * @return string|null
     */
    public function getMethod(): ?string
    {
        return $this->method;
    }

    /**
     * @param string|null $method
     */
    public function setMethod(?string $method): void
    {
        $this->method = $method;
    }

    /**
     * @return string|null
     */
    public function getRequestUri(): ?string
    {
        return $this->requestUri;
    }

    /**
     * @param string|null $requestUri
     */
    public function setRequestUri(?string $requestUri): void
    {
        $this->requestUri = $requestUri;
    }

    /**
     * @return array|null
     */
    public function getHeaders(): ?array
    {
        return $this->headers;
    }

    /**
     * @param array|null $headers
     */
    public function setHeaders(?array $headers): void
    {
        $this->headers = $headers;
    }

    /**
     * @return array|null
     */
    public function getCookies(): Cookies
    {
        return $this->cookies;
    }

    /**
     * @param array|null $cookies
     */
    public function setCookies(?array $cookies): void
    {
        $this->cookies = Cookies::create($cookies);
    }

    /**
     * @return string|null
     */
    public function getQueryString(): ?string
    {
        return $this->queryString;
    }

    /**
     * @param string|null $queryString
     */
    public function setQueryString(?string $queryString): void
    {
        $this->queryString = $queryString;
    }

    /**
     * @return string|null
     */
    public function getPostValues(): ?string
    {
        return $this->postValues;
    }

    /**
     * @param string|null $postValues
     */
    public function setPostValues(?string $postValues): void
    {
        $this->postValues = $postValues;
    }

    /**
     * @return array|null
     */
    public function getFormValues(): ?array
    {
        return $this->formValues;
    }

    /**
     * @param array|null $formValues
     */
    public function setFormValues(?array $formValues): void
    {
        $this->formValues = $formValues;
    }

    /**
     * @return array|null
     */
    public function getFormFiles(): ?array
    {
        return $this->formFiles;
    }

    /**
     * @param array|null $formFiles
     */
    public function setFormFiles(?array $formFiles): void
    {
        $this->formFiles = $formFiles;
    }

    /**
     * @return string|null
     */
    public function getRemoteAddr(): ?string
    {
        return $this->remoteAddr;
    }

    /**
     * @param string|null $remoteAddr
     */
    public function setRemoteAddr(?string $remoteAddr): void
    {
        $this->remoteAddr = $remoteAddr;
    }
}
