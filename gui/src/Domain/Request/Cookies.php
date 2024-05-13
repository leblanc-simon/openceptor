<?php

namespace App\Domain\Request;

use Symfony\Component\HttpFoundation\Cookie;

class Cookies
{
    private function __construct(public readonly array $datas)
    {
    }

    public static function create(?array $cookies)
    {
        if (null === $cookies) {
            $cookies = [];
        }

        $datas = [];

        foreach ($cookies as $cookie) {
            $datas[] = Cookie::fromString($cookie);
        }

        return new self($datas);
    }
}
