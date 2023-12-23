#ifndef LIBETAT_PDF_H
#define LIBETAT_PDF_H

#include "meta.h"

namespace lebetat {
    namespace pdf {
        class PdfObject {
            public:
                virtual void to_pdf();
        };

        class PdfText : public PdfObject {
            void to_pdf() override;
        };

        class PdfImg : public PdfObject {};

        class PdfPage : public PdfObject {};

        class PdfDocument : public PdfObject {};
    }//namespace pdf
}//namespace lebetat

#endif//LIBETAT_PDF_H