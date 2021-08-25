package client

import (
	"testing"
)

func TestDescClient_GetDesc(t *testing.T) {
	d := NewDetailClient(apiKey)
	t.Log(d.GetDescImgs("636828464656"))
}

func TestOptions(t *testing.T) {

}

var jsonStr = `{
    "ErrorCode": "Ok",
    "SubErrorCode": {},
    "RequestId": "86e5659a-3d4a-45ae-808f-8cc183345633",
    "RequestTime": 392.4139,
    "OtapiItemFullInfo": {
        "Id": "623394586347",
        "ErrorCode": "Ok",
        "HasError": false,
        "ProviderType": "Taobao",
        "UpdatedTime": "2021-08-25T15:15:15.291",
        "Title": "Vero Moda autumn retro washed one-shoulder waist A-line denim dress female | 319342505a",
        "OriginalTitle": "Vero Moda秋季复古水洗一字肩收腰A摆牛仔连衣裙女|319342505a",
        "CategoryId": "50010850",
        "ExternalCategoryId": "50010850",
        "VendorId": "veromoda官方outlets店",
        "VendorName": "veromoda官方outlets店",
        "VendorScore": 0,
        "BrandId": "ot:36071",
        "BrandName": "Vero Moda",
        "TaobaoItemUrl": "https://item.taobao.com/item.htm?id=623394586347",
        "ExternalItemUrl": "https://item.taobao.com/item.htm?id=623394586347",
        "MainPictureUrl": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2201198031066/O1CN01bzWuuh1JkGuElTWuH_!!2201198031066.jpg",
        "StuffStatus": "New",
        "Volume": 0,
        "Price": {
            "OriginalPrice": 699.0,
            "MarginPrice": 705.0,
            "OriginalCurrencyCode": "CNY",
            "ConvertedPriceList": {
                "Internal": {
                    "Price": 705.0,
                    "Sign": "元",
                    "Code": "CNY"
                },
                "DisplayedMoneys": [
                    {
                        "Price": 108.84,
                        "Sign": "$",
                        "Code": "USD"
                    }
                ]
            },
            "ConvertedPrice": "108.84$",
            "ConvertedPriceWithoutSign": "108.84",
            "CurrencySign": "$",
            "CurrencyName": "USD",
            "IsDeliverable": true,
            "DeliveryPrice": {
                "OriginalPrice": 6.0,
                "MarginPrice": 6.0,
                "OriginalCurrencyCode": "CNY",
                "ConvertedPriceList": {
                    "Internal": {
                        "Price": 6.0,
                        "Sign": "元",
                        "Code": "CNY"
                    },
                    "DisplayedMoneys": [
                        {
                            "Price": 0.93,
                            "Sign": "$",
                            "Code": "USD"
                        }
                    ]
                }
            },
            "OneItemDeliveryPrice": {
                "OriginalPrice": 6.0,
                "MarginPrice": 6.0,
                "OriginalCurrencyCode": "CNY",
                "ConvertedPriceList": {
                    "Internal": {
                        "Price": 6.0,
                        "Sign": "元",
                        "Code": "CNY"
                    },
                    "DisplayedMoneys": [
                        {
                            "Price": 0.93,
                            "Sign": "$",
                            "Code": "USD"
                        }
                    ]
                }
            },
            "PriceWithoutDelivery": {
                "OriginalPrice": 699.0,
                "MarginPrice": 699.0,
                "OriginalCurrencyCode": "CNY",
                "ConvertedPriceList": {
                    "Internal": {
                        "Price": 699.0,
                        "Sign": "元",
                        "Code": "CNY"
                    },
                    "DisplayedMoneys": [
                        {
                            "Price": 107.91,
                            "Sign": "$",
                            "Code": "USD"
                        }
                    ]
                }
            },
            "OneItemPriceWithoutDelivery": {
                "OriginalPrice": 699.0,
                "MarginPrice": 699.0,
                "OriginalCurrencyCode": "CNY",
                "ConvertedPriceList": {
                    "Internal": {
                        "Price": 699.0,
                        "Sign": "元",
                        "Code": "CNY"
                    },
                    "DisplayedMoneys": [
                        {
                            "Price": 107.92,
                            "Sign": "$",
                            "Code": "USD"
                        }
                    ]
                }
            }
        },
        "Pictures": [
            {
                "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2201198031066/O1CN01bzWuuh1JkGuElTWuH_!!2201198031066.jpg",
                "Small": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2201198031066/O1CN01bzWuuh1JkGuElTWuH_!!2201198031066.jpg_100x100q90.jpg",
                    "Width": 100,
                    "Height": 100
                },
                "Medium": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2201198031066/O1CN01bzWuuh1JkGuElTWuH_!!2201198031066.jpg_310x310q90.jpg",
                    "Width": 310,
                    "Height": 310
                },
                "Large": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2201198031066/O1CN01bzWuuh1JkGuElTWuH_!!2201198031066.jpg_600x600q90.jpg",
                    "Width": 600,
                    "Height": 600
                },
                "IsMain": true
            },
            {
                "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i1/2201198031066/O1CN01p0v41H1JkGp40btGr_!!2201198031066.jpg",
                "Small": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i1/2201198031066/O1CN01p0v41H1JkGp40btGr_!!2201198031066.jpg_100x100q90.jpg",
                    "Width": 100,
                    "Height": 100
                },
                "Medium": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i1/2201198031066/O1CN01p0v41H1JkGp40btGr_!!2201198031066.jpg_310x310q90.jpg",
                    "Width": 310,
                    "Height": 310
                },
                "Large": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i1/2201198031066/O1CN01p0v41H1JkGp40btGr_!!2201198031066.jpg_600x600q90.jpg",
                    "Width": 600,
                    "Height": 600
                },
                "IsMain": false
            },
            {
                "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i3/2201198031066/O1CN01GK0bxt1JkGp0rOXqq_!!2201198031066.jpg",
                "Small": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i3/2201198031066/O1CN01GK0bxt1JkGp0rOXqq_!!2201198031066.jpg_100x100q90.jpg",
                    "Width": 100,
                    "Height": 100
                },
                "Medium": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i3/2201198031066/O1CN01GK0bxt1JkGp0rOXqq_!!2201198031066.jpg_310x310q90.jpg",
                    "Width": 310,
                    "Height": 310
                },
                "Large": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i3/2201198031066/O1CN01GK0bxt1JkGp0rOXqq_!!2201198031066.jpg_600x600q90.jpg",
                    "Width": 600,
                    "Height": 600
                },
                "IsMain": false
            },
            {
                "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2201198031066/O1CN013Wo0Ye1JkGp7R8m2F_!!2201198031066.jpg",
                "Small": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2201198031066/O1CN013Wo0Ye1JkGp7R8m2F_!!2201198031066.jpg_100x100q90.jpg",
                    "Width": 100,
                    "Height": 100
                },
                "Medium": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2201198031066/O1CN013Wo0Ye1JkGp7R8m2F_!!2201198031066.jpg_310x310q90.jpg",
                    "Width": 310,
                    "Height": 310
                },
                "Large": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2201198031066/O1CN013Wo0Ye1JkGp7R8m2F_!!2201198031066.jpg_600x600q90.jpg",
                    "Width": 600,
                    "Height": 600
                },
                "IsMain": false
            },
            {
                "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i1/2201198031066/O1CN01bO1K4B1JkGp1Ls70Z_!!2201198031066.jpg",
                "Small": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i1/2201198031066/O1CN01bO1K4B1JkGp1Ls70Z_!!2201198031066.jpg_100x100q90.jpg",
                    "Width": 100,
                    "Height": 100
                },
                "Medium": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i1/2201198031066/O1CN01bO1K4B1JkGp1Ls70Z_!!2201198031066.jpg_310x310q90.jpg",
                    "Width": 310,
                    "Height": 310
                },
                "Large": {
                    "Url": "https://img.alicdn.com/imgextra///img.alicdn.com/imgextra/i1/2201198031066/O1CN01bO1K4B1JkGp1Ls70Z_!!2201198031066.jpg_600x600q90.jpg",
                    "Width": 600,
                    "Height": 600
                },
                "IsMain": false
            }
        ],
        "Location": {
            "State": "Tianjin"
        },
        "Features": [
            "Tmall"
        ],
        "FeaturedValues": [
            {
                "Name": "TaobaoVendorId",
                "Value": "2201198031066"
            },
            {
                "Name": "SalesInLast30Days",
                "Value": "91"
            },
            {
                "Name": "reviews",
                "Value": "35"
            }
        ],
        "IsSellAllowed": false,
        "SellDisallowReason": 5,
        "PhysicalParameters": {},
        "IsFiltered": false,
        "Videos": [
            {
                "PreviewUrl": "https://img.alicdn.com/imgextra/i2/2201198031066/O1CN01bzWuuh1JkGuElTWuH_!!2201198031066.jpg",
                "Url": "https://cloud.video.taobao.com/play/u/2201198031066/p/1/e/6/t/8/275592552146.mp4"
            }
        ],
        "HasInternalDelivery": true,
        "DeliveryCosts": [
            {
                "AreaCode": "110000",
                "Mode": "express",
                "Price": {
                    "OriginalPrice": 6.0,
                    "MarginPrice": 6.0,
                    "OriginalCurrencyCode": "CNY",
                    "ConvertedPriceList": {
                        "Internal": {
                            "Price": 6.0,
                            "Sign": "元",
                            "Code": "CNY"
                        },
                        "DisplayedMoneys": [
                            {
                                "Price": 0.93,
                                "Sign": "$",
                                "Code": "USD"
                            }
                        ]
                    },
                    "ConvertedPrice": "0.93$",
                    "ConvertedPriceWithoutSign": "0.93",
                    "CurrencySign": "$",
                    "CurrencyName": "USD",
                    "IsDeliverable": false,
                    "PriceWithoutDelivery": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    }
                },
                "StartCost": 6.0,
                "StartWeight": 0.0,
                "AddWeight": 0.0,
                "AddCost": 0.0
            }
        ],
        "Attributes": [
            {
                "Pid": "20509",
                "Vid": "397490054",
                "PropertyName": "Size",
                "Value": "155/76A/XS",
                "OriginalPropertyName": "尺码",
                "OriginalValue": "155/76A/XS",
                "IsConfigurator": true
            },
            {
                "Pid": "20509",
                "Vid": "387654415",
                "PropertyName": "Size",
                "Value": "160/80A/S",
                "OriginalPropertyName": "尺码",
                "OriginalValue": "160/80A/S",
                "IsConfigurator": true
            },
            {
                "Pid": "20509",
                "Vid": "387654416",
                "PropertyName": "Size",
                "Value": "165/84A/M",
                "OriginalPropertyName": "尺码",
                "OriginalValue": "165/84A/M",
                "IsConfigurator": true
            },
            {
                "Pid": "20509",
                "Vid": "387654417",
                "PropertyName": "Size",
                "Value": "170/88A/L",
                "OriginalPropertyName": "尺码",
                "OriginalValue": "170/88A/L",
                "IsConfigurator": true
            },
            {
                "Pid": "20509",
                "Vid": "39331866",
                "PropertyName": "Size",
                "Value": "175/92A/XL",
                "OriginalPropertyName": "尺码",
                "OriginalValue": "175/92A/XL",
                "IsConfigurator": true
            },
            {
                "Pid": "20509",
                "Vid": "427664426",
                "PropertyName": "Size",
                "Value": "180/96A/XXL",
                "OriginalPropertyName": "尺码",
                "OriginalValue": "180/96A/XXL",
                "IsConfigurator": true
            },
            {
                "Pid": "1627207",
                "Vid": "1631765841",
                "PropertyName": "Color classification",
                "Value": "J3E extra deep denim blue",
                "OriginalPropertyName": "颜色分类",
                "OriginalValue": "J3E特深牛仔蓝色",
                "IsConfigurator": true,
                "ImageUrl": "https://img.alicdn.com/imgextra/i4/2201198031066/O1CN01oA4wvP1JkGokO39Ut_!!2201198031066.jpg",
                "MiniImageUrl": "https://img.alicdn.com/imgextra/i4/2201198031066/O1CN01oA4wvP1JkGokO39Ut_!!2201198031066.jpg_100x100q90.jpg"
            },
            {
                "Pid": "品牌",
                "Vid": "Vero Moda",
                "PropertyName": "Brand",
                "Value": "Vero Moda",
                "OriginalPropertyName": "品牌",
                "OriginalValue": "Vero Moda",
                "IsConfigurator": false
            },
            {
                "Pid": "适用年龄",
                "Vid": "25-29周岁",
                "PropertyName": "Applicable age",
                "Value": "25-29 years old",
                "OriginalPropertyName": "适用年龄",
                "OriginalValue": "25-29周岁",
                "IsConfigurator": false
            },
            {
                "Pid": "图案",
                "Vid": "纯色",
                "PropertyName": "Pattern",
                "Value": "Solid color",
                "OriginalPropertyName": "图案",
                "OriginalValue": "纯色",
                "IsConfigurator": false
            },
            {
                "Pid": "风格",
                "Vid": "通勤",
                "PropertyName": "Style",
                "Value": "Commute",
                "OriginalPropertyName": "风格",
                "OriginalValue": "通勤",
                "IsConfigurator": false
            },
            {
                "Pid": "通勤",
                "Vid": "简约",
                "PropertyName": "Commute",
                "Value": "Simplicity",
                "OriginalPropertyName": "通勤",
                "OriginalValue": "简约",
                "IsConfigurator": false
            },
            {
                "Pid": "领型",
                "Vid": "一字领",
                "PropertyName": "Collar type",
                "Value": "One word collar",
                "OriginalPropertyName": "领型",
                "OriginalValue": "一字领",
                "IsConfigurator": false
            },
            {
                "Pid": "腰型",
                "Vid": "中腰",
                "PropertyName": "Waist type",
                "Value": "Mid waist",
                "OriginalPropertyName": "腰型",
                "OriginalValue": "中腰",
                "IsConfigurator": false
            },
            {
                "Pid": "衣门襟",
                "Vid": "拉链",
                "PropertyName": "Clothing placket",
                "Value": "Zipper",
                "OriginalPropertyName": "衣门襟",
                "OriginalValue": "拉链",
                "IsConfigurator": false
            },
            {
                "Pid": "袖型",
                "Vid": "常规",
                "PropertyName": "Sleeve type",
                "Value": "Regular",
                "OriginalPropertyName": "袖型",
                "OriginalValue": "常规",
                "IsConfigurator": false
            },
            {
                "Pid": "组合形式",
                "Vid": "单件",
                "PropertyName": "Combination form",
                "Value": "Single piece",
                "OriginalPropertyName": "组合形式",
                "OriginalValue": "单件",
                "IsConfigurator": false
            },
            {
                "Pid": "货号",
                "Vid": "319342505a",
                "PropertyName": "Item No.",
                "Value": "319342505a",
                "OriginalPropertyName": "货号",
                "OriginalValue": "319342505a",
                "IsConfigurator": false
            },
            {
                "Pid": "裙型",
                "Vid": "A字裙",
                "PropertyName": "Skirt type",
                "Value": "A-line skirt",
                "OriginalPropertyName": "裙型",
                "OriginalValue": "A字裙",
                "IsConfigurator": false
            },
            {
                "Pid": "年份季节",
                "Vid": "2019年秋季",
                "PropertyName": "Year season",
                "Value": "Fall 2019",
                "OriginalPropertyName": "年份季节",
                "OriginalValue": "2019年秋季",
                "IsConfigurator": false
            },
            {
                "Pid": "袖长",
                "Vid": "短袖",
                "PropertyName": "Sleeve length",
                "Value": "Short sleeve",
                "OriginalPropertyName": "袖长",
                "OriginalValue": "短袖",
                "IsConfigurator": false
            },
            {
                "Pid": "裙长",
                "Vid": "短裙",
                "PropertyName": "Skirt length",
                "Value": "Short skirt",
                "OriginalPropertyName": "裙长",
                "OriginalValue": "短裙",
                "IsConfigurator": false
            },
            {
                "Pid": "款式",
                "Vid": "其他/other",
                "PropertyName": "Style",
                "Value": "Other",
                "OriginalPropertyName": "款式",
                "OriginalValue": "其他/other",
                "IsConfigurator": false
            },
            {
                "Pid": "销售渠道类型",
                "Vid": "商场同款(线上线下都销售)",
                "PropertyName": "Sales channel type",
                "Value": "Same style in shopping malls (sold online and offline)",
                "OriginalPropertyName": "销售渠道类型",
                "OriginalValue": "商场同款(线上线下都销售)",
                "IsConfigurator": false
            },
            {
                "Pid": "材质成分",
                "Vid": "棉71.3% 聚酯纤维27.5% 聚氨酯弹性纤维(氨纶)1.2%",
                "PropertyName": "Material composition",
                "Value": "Cotton 71.3% polyester fiber 27.5% polyurethane elastic fiber (spandex) 1.2%",
                "OriginalPropertyName": "材质成分",
                "OriginalValue": "棉71.3% 聚酯纤维27.5% 聚氨酯弹性纤维(氨纶)1.2%",
                "IsConfigurator": false
            }
        ],
        "HasHierarchicalConfigurators": false,
        "ConfiguredItems": [
            {
                "Id": "4409256473146",
                "Quantity": 21,
                "SalesCount": 0,
                "Price": {
                    "OriginalPrice": 699.0,
                    "MarginPrice": 705.0,
                    "OriginalCurrencyCode": "CNY",
                    "ConvertedPriceList": {
                        "Internal": {
                            "Price": 705.0,
                            "Sign": "元",
                            "Code": "CNY"
                        },
                        "DisplayedMoneys": [
                            {
                                "Price": 108.84,
                                "Sign": "$",
                                "Code": "USD"
                            }
                        ]
                    },
                    "ConvertedPrice": "108.84$",
                    "ConvertedPriceWithoutSign": "108.84",
                    "CurrencySign": "$",
                    "CurrencyName": "USD",
                    "IsDeliverable": true,
                    "DeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemDeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "PriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.91,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemPriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.92,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    }
                },
                "Configurators": [
                    {
                        "Pid": "20509",
                        "Vid": "387654416"
                    },
                    {
                        "Pid": "1627207",
                        "Vid": "1631765841"
                    }
                ]
            },
            {
                "Id": "4409256473148",
                "Quantity": 0,
                "SalesCount": 0,
                "Price": {
                    "OriginalPrice": 699.0,
                    "MarginPrice": 705.0,
                    "OriginalCurrencyCode": "CNY",
                    "ConvertedPriceList": {
                        "Internal": {
                            "Price": 705.0,
                            "Sign": "元",
                            "Code": "CNY"
                        },
                        "DisplayedMoneys": [
                            {
                                "Price": 108.84,
                                "Sign": "$",
                                "Code": "USD"
                            }
                        ]
                    },
                    "ConvertedPrice": "108.84$",
                    "ConvertedPriceWithoutSign": "108.84",
                    "CurrencySign": "$",
                    "CurrencyName": "USD",
                    "IsDeliverable": true,
                    "DeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemDeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "PriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.91,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemPriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.92,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    }
                },
                "Configurators": [
                    {
                        "Pid": "20509",
                        "Vid": "39331866"
                    },
                    {
                        "Pid": "1627207",
                        "Vid": "1631765841"
                    }
                ]
            },
            {
                "Id": "4409256473147",
                "Quantity": 0,
                "SalesCount": 0,
                "Price": {
                    "OriginalPrice": 699.0,
                    "MarginPrice": 705.0,
                    "OriginalCurrencyCode": "CNY",
                    "ConvertedPriceList": {
                        "Internal": {
                            "Price": 705.0,
                            "Sign": "元",
                            "Code": "CNY"
                        },
                        "DisplayedMoneys": [
                            {
                                "Price": 108.84,
                                "Sign": "$",
                                "Code": "USD"
                            }
                        ]
                    },
                    "ConvertedPrice": "108.84$",
                    "ConvertedPriceWithoutSign": "108.84",
                    "CurrencySign": "$",
                    "CurrencyName": "USD",
                    "IsDeliverable": true,
                    "DeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemDeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "PriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.91,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemPriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.92,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    }
                },
                "Configurators": [
                    {
                        "Pid": "20509",
                        "Vid": "387654417"
                    },
                    {
                        "Pid": "1627207",
                        "Vid": "1631765841"
                    }
                ]
            },
            {
                "Id": "4409256473144",
                "Quantity": 97,
                "SalesCount": 0,
                "Price": {
                    "OriginalPrice": 699.0,
                    "MarginPrice": 705.0,
                    "OriginalCurrencyCode": "CNY",
                    "ConvertedPriceList": {
                        "Internal": {
                            "Price": 705.0,
                            "Sign": "元",
                            "Code": "CNY"
                        },
                        "DisplayedMoneys": [
                            {
                                "Price": 108.84,
                                "Sign": "$",
                                "Code": "USD"
                            }
                        ]
                    },
                    "ConvertedPrice": "108.84$",
                    "ConvertedPriceWithoutSign": "108.84",
                    "CurrencySign": "$",
                    "CurrencyName": "USD",
                    "IsDeliverable": true,
                    "DeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemDeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "PriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.91,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemPriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.92,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    }
                },
                "Configurators": [
                    {
                        "Pid": "20509",
                        "Vid": "397490054"
                    },
                    {
                        "Pid": "1627207",
                        "Vid": "1631765841"
                    }
                ]
            },
            {
                "Id": "4409256473145",
                "Quantity": 98,
                "SalesCount": 0,
                "Price": {
                    "OriginalPrice": 699.0,
                    "MarginPrice": 705.0,
                    "OriginalCurrencyCode": "CNY",
                    "ConvertedPriceList": {
                        "Internal": {
                            "Price": 705.0,
                            "Sign": "元",
                            "Code": "CNY"
                        },
                        "DisplayedMoneys": [
                            {
                                "Price": 108.84,
                                "Sign": "$",
                                "Code": "USD"
                            }
                        ]
                    },
                    "ConvertedPrice": "108.84$",
                    "ConvertedPriceWithoutSign": "108.84",
                    "CurrencySign": "$",
                    "CurrencyName": "USD",
                    "IsDeliverable": true,
                    "DeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemDeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "PriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.91,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemPriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.92,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    }
                },
                "Configurators": [
                    {
                        "Pid": "20509",
                        "Vid": "387654415"
                    },
                    {
                        "Pid": "1627207",
                        "Vid": "1631765841"
                    }
                ]
            },
            {
                "Id": "4409256473149",
                "Quantity": 0,
                "SalesCount": 0,
                "Price": {
                    "OriginalPrice": 699.0,
                    "MarginPrice": 705.0,
                    "OriginalCurrencyCode": "CNY",
                    "ConvertedPriceList": {
                        "Internal": {
                            "Price": 705.0,
                            "Sign": "元",
                            "Code": "CNY"
                        },
                        "DisplayedMoneys": [
                            {
                                "Price": 108.84,
                                "Sign": "$",
                                "Code": "USD"
                            }
                        ]
                    },
                    "ConvertedPrice": "108.84$",
                    "ConvertedPriceWithoutSign": "108.84",
                    "CurrencySign": "$",
                    "CurrencyName": "USD",
                    "IsDeliverable": true,
                    "DeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemDeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "PriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.91,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemPriceWithoutDelivery": {
                        "OriginalPrice": 699.0,
                        "MarginPrice": 699.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 699.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 107.92,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    }
                },
                "Configurators": [
                    {
                        "Pid": "20509",
                        "Vid": "427664426"
                    },
                    {
                        "Pid": "1627207",
                        "Vid": "1631765841"
                    }
                ]
            }
        ],
        "MasterQuantity": 216,
        "Promotions": [
            {
                "Id": "ot-d-promotion-623394586347",
                "ErrorCode": "Ok",
                "HasError": false,
                "ProviderType": "Taobao",
                "Name": "Promotion",
                "Price": {
                    "OriginalPrice": 244.65,
                    "MarginPrice": 250.65,
                    "OriginalCurrencyCode": "CNY",
                    "ConvertedPriceList": {
                        "Internal": {
                            "Price": 250.65,
                            "Sign": "元",
                            "Code": "CNY"
                        },
                        "DisplayedMoneys": [
                            {
                                "Price": 38.7,
                                "Sign": "$",
                                "Code": "USD"
                            }
                        ]
                    },
                    "ConvertedPrice": "38.70$",
                    "ConvertedPriceWithoutSign": "38.70",
                    "CurrencySign": "$",
                    "CurrencyName": "USD",
                    "IsDeliverable": true,
                    "DeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemDeliveryPrice": {
                        "OriginalPrice": 6.0,
                        "MarginPrice": 6.0,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 6.0,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 0.93,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "PriceWithoutDelivery": {
                        "OriginalPrice": 244.65,
                        "MarginPrice": 244.65,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 244.65,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 37.77,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    },
                    "OneItemPriceWithoutDelivery": {
                        "OriginalPrice": 244.65,
                        "MarginPrice": 244.65,
                        "OriginalCurrencyCode": "CNY",
                        "ConvertedPriceList": {
                            "Internal": {
                                "Price": 244.65,
                                "Sign": "元",
                                "Code": "CNY"
                            },
                            "DisplayedMoneys": [
                                {
                                    "Price": 37.77,
                                    "Sign": "$",
                                    "Code": "USD"
                                }
                            ]
                        }
                    }
                },
                "ConfiguredItems": [
                    {
                        "Price": {
                            "OriginalPrice": 244.65,
                            "MarginPrice": 250.65,
                            "OriginalCurrencyCode": "CNY",
                            "ConvertedPriceList": {
                                "Internal": {
                                    "Price": 250.65,
                                    "Sign": "元",
                                    "Code": "CNY"
                                },
                                "DisplayedMoneys": [
                                    {
                                        "Price": 38.7,
                                        "Sign": "$",
                                        "Code": "USD"
                                    }
                                ]
                            },
                            "ConvertedPrice": "38.70$",
                            "ConvertedPriceWithoutSign": "38.70",
                            "CurrencySign": "$",
                            "CurrencyName": "USD",
                            "IsDeliverable": true,
                            "DeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemDeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "PriceWithoutDelivery": {
                                "OriginalPrice": 244.65,
                                "MarginPrice": 244.65,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 244.65,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 37.77,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemPriceWithoutDelivery": {
                                "OriginalPrice": 244.65,
                                "MarginPrice": 244.65,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 244.65,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 37.77,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            }
                        },
                        "Id": "4409256473146",
                        "ErrorCode": "Ok",
                        "HasError": false
                    },
                    {
                        "Price": {
                            "OriginalPrice": 699.0,
                            "MarginPrice": 705.0,
                            "OriginalCurrencyCode": "CNY",
                            "ConvertedPriceList": {
                                "Internal": {
                                    "Price": 705.0,
                                    "Sign": "元",
                                    "Code": "CNY"
                                },
                                "DisplayedMoneys": [
                                    {
                                        "Price": 108.84,
                                        "Sign": "$",
                                        "Code": "USD"
                                    }
                                ]
                            },
                            "ConvertedPrice": "108.84$",
                            "ConvertedPriceWithoutSign": "108.84",
                            "CurrencySign": "$",
                            "CurrencyName": "USD",
                            "IsDeliverable": true,
                            "DeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemDeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "PriceWithoutDelivery": {
                                "OriginalPrice": 699.0,
                                "MarginPrice": 699.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 699.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 107.91,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemPriceWithoutDelivery": {
                                "OriginalPrice": 699.0,
                                "MarginPrice": 699.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 699.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 107.92,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            }
                        },
                        "Id": "4409256473148",
                        "ErrorCode": "Ok",
                        "HasError": false
                    },
                    {
                        "Price": {
                            "OriginalPrice": 699.0,
                            "MarginPrice": 705.0,
                            "OriginalCurrencyCode": "CNY",
                            "ConvertedPriceList": {
                                "Internal": {
                                    "Price": 705.0,
                                    "Sign": "元",
                                    "Code": "CNY"
                                },
                                "DisplayedMoneys": [
                                    {
                                        "Price": 108.84,
                                        "Sign": "$",
                                        "Code": "USD"
                                    }
                                ]
                            },
                            "ConvertedPrice": "108.84$",
                            "ConvertedPriceWithoutSign": "108.84",
                            "CurrencySign": "$",
                            "CurrencyName": "USD",
                            "IsDeliverable": true,
                            "DeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemDeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "PriceWithoutDelivery": {
                                "OriginalPrice": 699.0,
                                "MarginPrice": 699.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 699.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 107.91,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemPriceWithoutDelivery": {
                                "OriginalPrice": 699.0,
                                "MarginPrice": 699.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 699.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 107.92,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            }
                        },
                        "Id": "4409256473147",
                        "ErrorCode": "Ok",
                        "HasError": false
                    },
                    {
                        "Price": {
                            "OriginalPrice": 244.65,
                            "MarginPrice": 250.65,
                            "OriginalCurrencyCode": "CNY",
                            "ConvertedPriceList": {
                                "Internal": {
                                    "Price": 250.65,
                                    "Sign": "元",
                                    "Code": "CNY"
                                },
                                "DisplayedMoneys": [
                                    {
                                        "Price": 38.7,
                                        "Sign": "$",
                                        "Code": "USD"
                                    }
                                ]
                            },
                            "ConvertedPrice": "38.70$",
                            "ConvertedPriceWithoutSign": "38.70",
                            "CurrencySign": "$",
                            "CurrencyName": "USD",
                            "IsDeliverable": true,
                            "DeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemDeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "PriceWithoutDelivery": {
                                "OriginalPrice": 244.65,
                                "MarginPrice": 244.65,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 244.65,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 37.77,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemPriceWithoutDelivery": {
                                "OriginalPrice": 244.65,
                                "MarginPrice": 244.65,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 244.65,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 37.77,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            }
                        },
                        "Id": "4409256473144",
                        "ErrorCode": "Ok",
                        "HasError": false
                    },
                    {
                        "Price": {
                            "OriginalPrice": 244.65,
                            "MarginPrice": 250.65,
                            "OriginalCurrencyCode": "CNY",
                            "ConvertedPriceList": {
                                "Internal": {
                                    "Price": 250.65,
                                    "Sign": "元",
                                    "Code": "CNY"
                                },
                                "DisplayedMoneys": [
                                    {
                                        "Price": 38.7,
                                        "Sign": "$",
                                        "Code": "USD"
                                    }
                                ]
                            },
                            "ConvertedPrice": "38.70$",
                            "ConvertedPriceWithoutSign": "38.70",
                            "CurrencySign": "$",
                            "CurrencyName": "USD",
                            "IsDeliverable": true,
                            "DeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemDeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "PriceWithoutDelivery": {
                                "OriginalPrice": 244.65,
                                "MarginPrice": 244.65,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 244.65,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 37.77,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemPriceWithoutDelivery": {
                                "OriginalPrice": 244.65,
                                "MarginPrice": 244.65,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 244.65,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 37.77,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            }
                        },
                        "Id": "4409256473145",
                        "ErrorCode": "Ok",
                        "HasError": false
                    },
                    {
                        "Price": {
                            "OriginalPrice": 699.0,
                            "MarginPrice": 705.0,
                            "OriginalCurrencyCode": "CNY",
                            "ConvertedPriceList": {
                                "Internal": {
                                    "Price": 705.0,
                                    "Sign": "元",
                                    "Code": "CNY"
                                },
                                "DisplayedMoneys": [
                                    {
                                        "Price": 108.84,
                                        "Sign": "$",
                                        "Code": "USD"
                                    }
                                ]
                            },
                            "ConvertedPrice": "108.84$",
                            "ConvertedPriceWithoutSign": "108.84",
                            "CurrencySign": "$",
                            "CurrencyName": "USD",
                            "IsDeliverable": true,
                            "DeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemDeliveryPrice": {
                                "OriginalPrice": 6.0,
                                "MarginPrice": 6.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 6.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 0.93,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "PriceWithoutDelivery": {
                                "OriginalPrice": 699.0,
                                "MarginPrice": 699.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 699.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 107.91,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            },
                            "OneItemPriceWithoutDelivery": {
                                "OriginalPrice": 699.0,
                                "MarginPrice": 699.0,
                                "OriginalCurrencyCode": "CNY",
                                "ConvertedPriceList": {
                                    "Internal": {
                                        "Price": 699.0,
                                        "Sign": "元",
                                        "Code": "CNY"
                                    },
                                    "DisplayedMoneys": [
                                        {
                                            "Price": 107.92,
                                            "Sign": "$",
                                            "Code": "USD"
                                        }
                                    ]
                                }
                            }
                        },
                        "Id": "4409256473149",
                        "ErrorCode": "Ok",
                        "HasError": false
                    }
                ]
            }
        ],
        "FirstLotQuantity": 1,
        "NextLotQuantity": 1,
        "WeightInfos": [],
        "ActualWeightInfo": {
            "Type": "Default",
            "DisplayName": "Approx weight",
            "Weight": 0.0
        }
    }
}`
